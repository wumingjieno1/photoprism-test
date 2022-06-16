package workers

import (
	"path"
	"path/filepath"
	"time"

	"github.com/photoprism/photoprism/internal/entity"
	"github.com/photoprism/photoprism/internal/event"
	"github.com/photoprism/photoprism/internal/mutex"
	"github.com/photoprism/photoprism/internal/photoprism"
	"github.com/photoprism/photoprism/internal/query"
	"github.com/photoprism/photoprism/internal/remote/webdav"
	"github.com/photoprism/photoprism/pkg/clean"
)

// Uploads local files to a remote account
func (worker *Sync) upload(a entity.Account) (complete bool, err error) {
	maxResults := 250

	// Get upload file list from database
	files, err := query.AccountUploads(a, maxResults)

	if err != nil {
		return false, err
	}

	if len(files) == 0 {
		log.Infof("sync: upload complete for %s", a.AccName)
		event.Publish("sync.uploaded", event.Data{"account": a})
		return true, nil
	}

	client := webdav.New(a.AccURL, a.AccUser, a.AccPass, webdav.Timeout(a.AccTimeout))
	existingDirs := make(map[string]string)

	for _, file := range files {
		if mutex.SyncWorker.Canceled() {
			return false, nil
		}

		fileName := photoprism.FileName(file.FileRoot, file.FileName)
		remoteName := path.Join(a.SyncPath, file.FileName)
		remoteDir := filepath.Dir(remoteName)

		if _, ok := existingDirs[remoteDir]; !ok {
			if err := client.CreateDir(remoteDir); err != nil {
				log.Errorf("sync: failed creating remote folder %s", remoteDir)
				continue // try again next time
			}
		}

		if err := client.Upload(fileName, remoteName); err != nil {
			worker.logError(err)
			continue // try again next time
		}

		log.Infof("sync: uploaded %s to %s (%s)", clean.Log(file.FileName), clean.Log(remoteName), a.AccName)

		fileSync := entity.NewFileSync(a.ID, remoteName)
		fileSync.Status = entity.FileSyncUploaded
		fileSync.RemoteDate = time.Now()
		fileSync.RemoteSize = file.FileSize
		fileSync.FileID = file.ID
		fileSync.Error = ""
		fileSync.Errors = 0

		if mutex.SyncWorker.Canceled() {
			return false, nil
		}

		worker.logError(entity.Db().Save(&fileSync).Error)
	}

	return false, nil
}
