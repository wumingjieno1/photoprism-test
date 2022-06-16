package workers

import (
	"fmt"
	"runtime/debug"
	"time"

	"github.com/photoprism/photoprism/internal/config"
	"github.com/photoprism/photoprism/internal/entity"
	"github.com/photoprism/photoprism/internal/event"
	"github.com/photoprism/photoprism/internal/form"
	"github.com/photoprism/photoprism/internal/mutex"
	"github.com/photoprism/photoprism/internal/remote"
	"github.com/photoprism/photoprism/internal/search"
)

// Sync represents a sync worker.
type Sync struct {
	conf *config.Config
}

// NewSync returns a new sync worker.
func NewSync(conf *config.Config) *Sync {
	return &Sync{
		conf: conf,
	}
}

// logError logs an error message if err is not nil.
func (worker *Sync) logError(err error) {
	if err != nil {
		log.Errorf("sync: %s", err.Error())
	}
}

// logWarn logs a warning message if err is not nil.
func (worker *Sync) logWarn(err error) {
	if err != nil {
		log.Warnf("sync: %s", err.Error())
	}
}

// Start starts the sync worker.
func (worker *Sync) Start() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("sync: %s (panic)\nstack: %s", r, debug.Stack())
			log.Error(err)
		}
	}()

	if err := mutex.SyncWorker.Start(); err != nil {
		return err
	}

	defer mutex.SyncWorker.Stop()

	f := form.SearchAccounts{
		Sync: true,
	}

	accounts, err := search.Accounts(f)

	for _, a := range accounts {
		if a.AccType != remote.ServiceWebDAV {
			continue
		}

		// Failed too often?
		if a.RetryLimit > 0 && a.AccErrors > a.RetryLimit {
			a.AccSync = false

			if err := entity.Db().Save(&a).Error; err != nil {
				worker.logError(err)
			} else {
				log.Warnf("sync: disabled sync, %s failed more than %d times", a.AccName, a.RetryLimit)
			}

			continue
		}

		// Values updated in account: AccError, AccErrors, SyncStatus, SyncDate
		accError := a.AccError
		accErrors := a.AccErrors
		syncStatus := a.SyncStatus
		syncDate := a.SyncDate
		synced := false

		switch a.SyncStatus {
		case entity.AccountSyncStatusRefresh:
			if complete, err := worker.refresh(a); err != nil {
				accErrors++
				accError = err.Error()
			} else if complete {
				accErrors = 0
				accError = ""

				if a.SyncDownload {
					syncStatus = entity.AccountSyncStatusDownload
				} else if a.SyncUpload {
					syncStatus = entity.AccountSyncStatusUpload
				} else {
					syncStatus = entity.AccountSyncStatusSynced
					syncDate.Time = time.Now()
					syncDate.Valid = true
				}
			}
		case entity.AccountSyncStatusDownload:
			if complete, err := worker.download(a); err != nil {
				accErrors++
				accError = err.Error()
				syncStatus = entity.AccountSyncStatusRefresh
			} else if complete {
				if a.SyncUpload {
					syncStatus = entity.AccountSyncStatusUpload
				} else {
					synced = true
					syncStatus = entity.AccountSyncStatusSynced
					syncDate.Time = time.Now()
					syncDate.Valid = true
				}
			}
		case entity.AccountSyncStatusUpload:
			if complete, err := worker.upload(a); err != nil {
				accErrors++
				accError = err.Error()
				syncStatus = entity.AccountSyncStatusRefresh
			} else if complete {
				synced = true
				syncStatus = entity.AccountSyncStatusSynced
				syncDate.Time = time.Now()
				syncDate.Valid = true
			}
		case entity.AccountSyncStatusSynced:
			if a.SyncDate.Valid && a.SyncDate.Time.Before(time.Now().Add(time.Duration(-1*a.SyncInterval)*time.Second)) {
				syncStatus = entity.AccountSyncStatusRefresh
			}
		default:
			syncStatus = entity.AccountSyncStatusRefresh
		}

		if mutex.SyncWorker.Canceled() {
			return nil
		}

		// Only update the following fields to avoid overwriting other settings
		if err := a.Updates(map[string]interface{}{
			"AccError":   accError,
			"AccErrors":  accErrors,
			"SyncStatus": syncStatus,
			"SyncDate":   syncDate}); err != nil {
			worker.logError(err)
		} else if synced {
			event.Publish("sync.synced", event.Data{"account": a})
		}
	}

	return err
}
