package api

import (
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/photoprism/photoprism/internal/acl"
	"github.com/photoprism/photoprism/internal/entity"
	"github.com/photoprism/photoprism/internal/event"
	"github.com/photoprism/photoprism/internal/form"
	"github.com/photoprism/photoprism/internal/i18n"
	"github.com/photoprism/photoprism/internal/photoprism"
	"github.com/photoprism/photoprism/internal/service"
	"github.com/photoprism/photoprism/pkg/clean"
	"github.com/photoprism/photoprism/pkg/txt"
)

// StartIndexing indexes media files in the "originals" folder.
//
// POST /api/v1/index
func StartIndexing(router *gin.RouterGroup) {
	router.POST("/index", func(c *gin.Context) {
		s := Auth(SessionID(c), acl.ResourcePhotos, acl.ActionUpdate)

		if s.Invalid() {
			AbortUnauthorized(c)
			return
		}

		conf := service.Config()
		settings := conf.Settings()

		if !settings.Features.Library {
			AbortFeatureDisabled(c)
			return
		}

		start := time.Now()

		var f form.IndexOptions

		if err := c.BindJSON(&f); err != nil {
			AbortBadRequest(c)
			return
		}

		// Configure index options.
		path := conf.OriginalsPath()
		convert := settings.Index.Convert && conf.SidecarWritable()
		skipArchived := settings.Index.SkipArchived

		indOpt := photoprism.NewIndexOptions(filepath.Clean(f.Path), f.Rescan, convert, true, false, skipArchived)

		if len(indOpt.Path) > 1 {
			event.InfoMsg(i18n.MsgIndexingFiles, clean.Log(indOpt.Path))
		} else {
			event.InfoMsg(i18n.MsgIndexingOriginals)
		}

		// Start indexing.
		ind := service.Index()
		indexed := ind.Start(indOpt)

		RemoveFromFolderCache(entity.RootOriginals)

		// Configure purge options.
		prgOpt := photoprism.PurgeOptions{
			Path:   filepath.Clean(f.Path),
			Ignore: indexed,
		}

		// Start purging.
		prg := service.Purge()

		if files, photos, err := prg.Start(prgOpt); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": txt.UpperFirst(err.Error())})
			return
		} else if len(files) > 0 || len(photos) > 0 {
			event.InfoMsg(i18n.MsgRemovedFilesAndPhotos, len(files), len(photos))
		}

		event.Publish("index.updating", event.Data{
			"step": "moments",
		})

		moments := service.Moments()

		if err := moments.Start(); err != nil {
			log.Warnf("moments: %s", err)
		}

		elapsed := int(time.Since(start).Seconds())

		msg := i18n.Msg(i18n.MsgIndexingCompletedIn, elapsed)

		event.Success(msg)
		event.Publish("index.completed", event.Data{"path": path, "seconds": elapsed})

		UpdateClientConfig()

		c.JSON(http.StatusOK, i18n.Response{Code: http.StatusOK, Msg: msg})
	})
}

// CancelIndexing stops indexing media files in the "originals" folder.
//
// DELETE /api/v1/index
func CancelIndexing(router *gin.RouterGroup) {
	router.DELETE("/index", func(c *gin.Context) {
		s := Auth(SessionID(c), acl.ResourcePhotos, acl.ActionUpdate)

		if s.Invalid() {
			AbortUnauthorized(c)
			return
		}

		conf := service.Config()

		if !conf.Settings().Features.Library {
			AbortFeatureDisabled(c)
			return
		}

		ind := service.Index()

		ind.Cancel()

		c.JSON(http.StatusOK, i18n.NewResponse(http.StatusOK, i18n.MsgIndexingCanceled))
	})
}
