package workers

import (
	"fmt"
	"os"

	"github.com/photoprism/photoprism/internal/entity"
	"github.com/photoprism/photoprism/internal/event"
	"github.com/photoprism/photoprism/internal/mutex"
	"github.com/photoprism/photoprism/internal/photoprism"
	"github.com/photoprism/photoprism/internal/query"
	"github.com/photoprism/photoprism/internal/remote/webdav"
	"github.com/photoprism/photoprism/internal/service"
	"github.com/photoprism/photoprism/pkg/fs"
)

type Downloads map[string][]entity.FileSync

// downloadPath returns a temporary download path.
func (worker *Sync) downloadPath() string {
	return worker.conf.TempPath() + "/sync"
}

// relatedDownloads returns files to be downloaded grouped by prefix.
func (worker *Sync) relatedDownloads(a entity.Account) (result Downloads, err error) {
	result = make(Downloads)
	maxResults := 1000

	// Get remote files from database
	files, err := query.FileSyncs(a.ID, entity.FileSyncNew, maxResults)

	if err != nil {
		return result, err
	}

	// Group results by directory and base name
	for i, file := range files {
		k := fs.AbsPrefix(file.RemoteName, worker.conf.Settings().StackSequences())

		result[k] = append(result[k], file)

		// Skip last 50 to make sure we see all related files
		if i > (maxResults - 50) {
			return result, nil
		}
	}

	return result, nil
}

// Downloads remote files in batches and imports / indexes them
func (worker *Sync) download(a entity.Account) (complete bool, err error) {
	// Set up index worker
	indexJobs := make(chan photoprism.IndexJob)

	go photoprism.IndexWorker(indexJobs)
	defer close(indexJobs)

	// Set up import worker
	importJobs := make(chan photoprism.ImportJob)
	go photoprism.ImportWorker(importJobs)
	defer close(importJobs)

	relatedFiles, err := worker.relatedDownloads(a)

	if err != nil {
		worker.logError(err)
		return false, err
	}

	if len(relatedFiles) == 0 {
		log.Infof("sync: download complete for %s", a.AccName)
		event.Publish("sync.downloaded", event.Data{"account": a})
		return true, nil
	}

	log.Infof("sync: downloading from %s", a.AccName)

	client := webdav.New(a.AccURL, a.AccUser, a.AccPass, webdav.Timeout(a.AccTimeout))

	var baseDir string

	if a.SyncFilenames {
		baseDir = worker.conf.OriginalsPath()
	} else {
		baseDir = fmt.Sprintf("%s/%d", worker.downloadPath(), a.ID)
	}

	done := make(map[string]bool)

	for _, files := range relatedFiles {
		for i, file := range files {
			if mutex.SyncWorker.Canceled() {
				return false, nil
			}

			// Failed too often?
			if a.RetryLimit > 0 && file.Errors > a.RetryLimit {
				log.Debugf("sync: downloading %s failed more than %d times", file.RemoteName, a.RetryLimit)
				continue
			}

			localName := baseDir + file.RemoteName

			if _, err := os.Stat(localName); err == nil {
				log.Warnf("sync: download skipped, %s already exists", localName)
				file.Status = entity.FileSyncExists
				file.Error = ""
				file.Errors = 0
			} else {
				if err := client.Download(file.RemoteName, localName, false); err != nil {
					file.Errors++
					file.Error = err.Error()
				} else {
					log.Infof("sync: downloaded %s from %s", file.RemoteName, a.AccName)
					file.Status = entity.FileSyncDownloaded
					file.Error = ""
					file.Errors = 0
				}

				if mutex.SyncWorker.Canceled() {
					return false, nil
				}
			}

			if err := entity.Db().Save(&file).Error; err != nil {
				worker.logError(err)
			} else {
				files[i] = file
			}
		}

		for _, file := range files {
			if file.Status != entity.FileSyncDownloaded {
				continue
			}

			mf, err := photoprism.NewMediaFile(baseDir + file.RemoteName)

			if err != nil || !mf.IsMedia() {
				continue
			}

			related, err := mf.RelatedFiles(worker.conf.Settings().StackSequences())

			if err != nil {
				worker.logWarn(err)
				continue
			}

			var rf photoprism.MediaFiles

			for _, f := range related.Files {
				if done[f.FileName()] {
					continue
				}

				rf = append(rf, f)
				done[f.FileName()] = true
			}

			done[mf.FileName()] = true
			related.Files = rf

			if a.SyncFilenames {
				log.Infof("sync: indexing %s and related files", file.RemoteName)
				indexJobs <- photoprism.IndexJob{
					FileName: mf.FileName(),
					Related:  related,
					IndexOpt: photoprism.IndexOptionsAll(),
					Ind:      service.Index(),
				}
			} else {
				log.Infof("sync: importing %s and related files", file.RemoteName)
				importJobs <- photoprism.ImportJob{
					FileName:  mf.FileName(),
					Related:   related,
					IndexOpt:  photoprism.IndexOptionsAll(),
					ImportOpt: photoprism.ImportOptionsMove(baseDir),
					Imp:       service.Import(),
				}
			}
		}
	}

	// Any files downloaded?
	if len(done) > 0 {
		// Update precalculated photo and file counts.
		worker.logWarn(entity.UpdateCounts())

		// Update album, subject, and label cover thumbs.
		worker.logWarn(query.UpdateCovers())
	}

	return false, nil
}
