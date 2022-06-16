package photoprism

import (
	"errors"
	"fmt"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"sync"

	"github.com/karrick/godirwalk"

	"github.com/photoprism/photoprism/internal/classify"
	"github.com/photoprism/photoprism/internal/config"
	"github.com/photoprism/photoprism/internal/entity"
	"github.com/photoprism/photoprism/internal/event"
	"github.com/photoprism/photoprism/internal/face"
	"github.com/photoprism/photoprism/internal/mutex"
	"github.com/photoprism/photoprism/internal/nsfw"
	"github.com/photoprism/photoprism/pkg/clean"
	"github.com/photoprism/photoprism/pkg/fs"
	"github.com/photoprism/photoprism/pkg/media"
)

// Index represents an indexer that indexes files in the originals directory.
type Index struct {
	conf         *config.Config
	tensorFlow   *classify.TensorFlow
	nsfwDetector *nsfw.Detector
	faceNet      *face.Net
	convert      *Convert
	files        *Files
	photos       *Photos
	findFaces    bool
	findLabels   bool
}

// NewIndex returns a new indexer and expects its dependencies as arguments.
func NewIndex(conf *config.Config, tensorFlow *classify.TensorFlow, nsfwDetector *nsfw.Detector, faceNet *face.Net, convert *Convert, files *Files, photos *Photos) *Index {
	if conf == nil {
		log.Errorf("index: config is nil")
		return nil
	}

	i := &Index{
		conf:         conf,
		tensorFlow:   tensorFlow,
		nsfwDetector: nsfwDetector,
		faceNet:      faceNet,
		convert:      convert,
		files:        files,
		photos:       photos,
		findFaces:    !conf.DisableFaces(),
		findLabels:   !conf.DisableClassification(),
	}

	return i
}

func (ind *Index) originalsPath() string {
	return ind.conf.OriginalsPath()
}

func (ind *Index) thumbPath() string {
	return ind.conf.ThumbCachePath()
}

// Cancel stops the current indexing operation.
func (ind *Index) Cancel() {
	mutex.MainWorker.Cancel()
}

// Start indexes media files in the "originals" folder.
func (ind *Index) Start(o IndexOptions) fs.Done {
	defer func() {
		if r := recover(); r != nil {
			log.Errorf("index: %s (panic)\nstack: %s", r, debug.Stack())
		}
	}()

	done := make(fs.Done)

	if ind.conf == nil {
		log.Errorf("index: config is nil")
		return done
	}

	originalsPath := ind.originalsPath()
	optionsPath := filepath.Join(originalsPath, o.Path)

	if !fs.PathExists(optionsPath) {
		event.Error(fmt.Sprintf("index: %s does not exist", clean.Log(optionsPath)))
		return done
	}

	if err := mutex.MainWorker.Start(); err != nil {
		event.Error(fmt.Sprintf("index: %s", err.Error()))
		return done
	}

	defer mutex.MainWorker.Stop()

	if err := ind.tensorFlow.Init(); err != nil {
		log.Errorf("index: %s", err.Error())

		return done
	}

	jobs := make(chan IndexJob)

	// Start a fixed number of goroutines to index files.
	var wg sync.WaitGroup
	var numWorkers = ind.conf.Workers()
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func() {
			IndexWorker(jobs) // HLc
			wg.Done()
		}()
	}

	if err := ind.files.Init(); err != nil {
		log.Errorf("index: %s", err)
	}

	defer ind.files.Done()

	filesIndexed := 0
	skipRaw := ind.conf.DisableRaw()
	ignore := fs.NewIgnoreList(fs.IgnoreFile, true, false)

	if err := ignore.Dir(originalsPath); err != nil {
		log.Infof("index: %s", err)
	}

	ignore.Log = func(fileName string) {
		log.Infof(`index: ignored "%s"`, fs.RelName(fileName, originalsPath))
	}

	err := godirwalk.Walk(optionsPath, &godirwalk.Options{
		ErrorCallback: func(fileName string, err error) godirwalk.ErrorAction {
			return godirwalk.SkipNode
		},
		Callback: func(fileName string, info *godirwalk.Dirent) error {
			defer func() {
				if r := recover(); r != nil {
					log.Errorf("index: %s (panic)\nstack: %s", r, debug.Stack())
				}
			}()

			if mutex.MainWorker.Canceled() {
				return errors.New("canceled")
			}

			isDir := info.IsDir()
			isSymlink := info.IsSymlink()
			relName := fs.RelName(fileName, originalsPath)

			if skip, result := fs.SkipWalk(fileName, isDir, isSymlink, done, ignore); skip {
				if (isSymlink || isDir) && result != filepath.SkipDir {
					folder := entity.NewFolder(entity.RootOriginals, relName, fs.BirthTime(fileName))

					if err := folder.Create(); err == nil {
						log.Infof("index: added folder /%s", folder.Path)
					}
				}

				if isDir {
					event.Publish("index.folder", event.Data{
						"filePath": relName,
					})
				}

				return result
			}

			done[fileName] = fs.Found

			if !media.MainFile(fileName) {
				return nil
			}

			mf, err := NewMediaFile(fileName)

			// Check if file exists and is not empty.
			if err != nil {
				log.Warnf("index: %s", err)
				return nil
			}

			// Ignore RAW images?
			if mf.IsRaw() && skipRaw {
				log.Infof("index: skipped raw %s", clean.Log(mf.RootRelName()))
				return nil
			}

			// Skip?
			if ind.files.Indexed(relName, entity.RootOriginals, mf.modTime, o.Rescan) {
				return nil
			}

			// Find related files to index.
			related, err := mf.RelatedFiles(ind.conf.Settings().StackSequences())

			if err != nil {
				log.Warnf("index: %s", err.Error())
				return nil
			}

			var files MediaFiles

			for _, f := range related.Files {
				if done[f.FileName()].Processed() {
					continue
				}

				if f.FileSize() == 0 || ind.files.Indexed(f.RootRelName(), f.Root(), f.ModTime(), o.Rescan) {
					done[f.FileName()] = fs.Found
					continue
				}

				files = append(files, f)
				filesIndexed++
				done[f.FileName()] = fs.Processed
			}

			done[fileName] = fs.Processed

			if len(files) == 0 || related.Main == nil {
				// Nothing to do.
				return nil
			}

			related.Files = files

			jobs <- IndexJob{
				FileName: mf.FileName(),
				Related:  related,
				IndexOpt: o,
				Ind:      ind,
			}

			return nil
		},
		Unsorted:            false,
		FollowSymbolicLinks: true,
	})

	close(jobs)
	wg.Wait()

	if err != nil {
		log.Error(err.Error())
	}

	if filesIndexed > 0 {
		event.Publish("index.updating", event.Data{
			"step": "faces",
		})

		// Run facial recognition if enabled.
		if w := NewFaces(ind.conf); w.Disabled() {
			log.Debugf("index: skipping facial recognition")
		} else if err := w.Start(FacesOptionsDefault()); err != nil {
			log.Errorf("index: %s", err)
		}

		event.Publish("index.updating", event.Data{
			"step": "counts",
		})

		// Update precalculated photo and file counts.
		if err := entity.UpdateCounts(); err != nil {
			log.Warnf("index: %s (update counts)", err)
		}
	} else {
		log.Infof("index: found no new or modified files")
	}

	runtime.GC()

	return done
}

// FileName indexes a single file and returns the result.
func (ind *Index) FileName(fileName string, o IndexOptions) (result IndexResult) {
	file, err := NewMediaFile(fileName)

	if err != nil {
		result.Err = err
		result.Status = IndexFailed

		return result
	}

	related, err := file.RelatedFiles(false)

	if err != nil {
		result.Err = err
		result.Status = IndexFailed

		return result
	}

	return IndexRelated(related, ind, o)
}
