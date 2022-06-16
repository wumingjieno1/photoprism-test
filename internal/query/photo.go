package query

import (
	"time"

	"github.com/dustin/go-humanize/english"
	"github.com/jinzhu/gorm"

	"github.com/photoprism/photoprism/internal/entity"
	"github.com/photoprism/photoprism/internal/mutex"
)

// PhotoByID returns a Photo based on the ID.
func PhotoByID(photoID uint64) (photo entity.Photo, err error) {
	if err := UnscopedDb().Where("id = ?", photoID).
		Preload("Labels", func(db *gorm.DB) *gorm.DB {
			return db.Order("photos_labels.uncertainty ASC, photos_labels.label_id DESC")
		}).
		Preload("Labels.Label").
		Preload("Camera").
		Preload("Lens").
		Preload("Details").
		Preload("Place").
		Preload("Cell").
		Preload("Cell.Place").
		First(&photo).Error; err != nil {
		return photo, err
	}

	return photo, nil
}

// PhotoByUID returns a Photo based on the UID.
func PhotoByUID(photoUID string) (photo entity.Photo, err error) {
	if err := UnscopedDb().Where("photo_uid = ?", photoUID).
		Preload("Labels", func(db *gorm.DB) *gorm.DB {
			return db.Order("photos_labels.uncertainty ASC, photos_labels.label_id DESC")
		}).
		Preload("Labels.Label").
		Preload("Camera").
		Preload("Lens").
		Preload("Details").
		Preload("Place").
		Preload("Cell").
		Preload("Cell.Place").
		First(&photo).Error; err != nil {
		return photo, err
	}

	return photo, nil
}

// PhotoPreloadByUID returns a Photo based on the UID with all dependencies preloaded.
func PhotoPreloadByUID(photoUID string) (photo entity.Photo, err error) {
	if err := UnscopedDb().Where("photo_uid = ?", photoUID).
		Preload("Labels", func(db *gorm.DB) *gorm.DB {
			return db.Order("photos_labels.uncertainty ASC, photos_labels.label_id DESC")
		}).
		Preload("Labels.Label").
		Preload("Camera").
		Preload("Lens").
		Preload("Details").
		Preload("Place").
		Preload("Cell").
		Preload("Cell.Place").
		First(&photo).Error; err != nil {
		return photo, err
	}

	photo.PreloadMany()

	return photo, nil
}

// PhotosMissing returns photo entities without existing files.
func PhotosMissing(limit int, offset int) (entities entity.Photos, err error) {
	err = Db().
		Select("photos.*").
		Where("id NOT IN (SELECT photo_id FROM files WHERE file_missing = 0 AND file_root = '/' AND deleted_at IS NULL)").
		Where("photos.photo_type <> ?", entity.MediaText).
		Group("photos.id").
		Limit(limit).Offset(offset).Find(&entities).Error

	return entities, err
}

// PhotosMetadataUpdate returns photos selected for metadata maintenance.
func PhotosMetadataUpdate(limit, offset int, delay, interval time.Duration) (entities entity.Photos, err error) {
	err = Db().
		Preload("Labels", func(db *gorm.DB) *gorm.DB {
			return db.Order("photos_labels.uncertainty ASC, photos_labels.label_id DESC")
		}).
		Preload("Labels.Label").
		Preload("Camera").
		Preload("Lens").
		Preload("Details").
		Preload("Place").
		Preload("Cell").
		Preload("Cell.Place").
		Where("checked_at IS NULL OR checked_at < ?", time.Now().Add(-1*interval)).
		Where("updated_at < ? OR (cell_id = 'zz' AND photo_lat <> 0)", time.Now().Add(-1*delay)).
		Order("photos.ID ASC").Limit(limit).Offset(offset).Find(&entities).Error

	return entities, err
}

// OrphanPhotos finds orphan index entries that may be removed.
func OrphanPhotos() (photos entity.Photos, err error) {
	err = UnscopedDb().
		Raw(`SELECT * FROM photos WHERE 
			deleted_at IS NOT NULL 
			AND photo_quality = -1 
			AND id NOT IN (SELECT photo_id FROM files WHERE files.deleted_at IS NULL)`).
		Find(&photos).Error

	return photos, err
}

// FixPrimaries tries to set a primary file for photos that have none.
func FixPrimaries() error {
	mutex.Index.Lock()
	defer mutex.Index.Unlock()

	start := time.Now()

	var photos entity.Photos

	// Remove primary file flag from broken or missing files.
	if err := UnscopedDb().Table(entity.File{}.TableName()).
		Where("file_error <> '' OR file_missing = 1").
		UpdateColumn("file_primary", 0).Error; err != nil {
		return err
	}

	// Find photos without primary file.
	if err := UnscopedDb().
		Raw(`SELECT * FROM photos 
			WHERE deleted_at IS NULL 
			AND id NOT IN (SELECT photo_id FROM files WHERE file_primary = 1)`).
		Find(&photos).Error; err != nil {
		return err
	}

	if len(photos) == 0 {
		log.Debugf("index: found no photos without primary file [%s]", time.Since(start))
		return nil
	}

	// Try to find matching primary files.
	for _, p := range photos {
		log.Debugf("index: searching primary file for %s", p.PhotoUID)

		if err := p.SetPrimary(""); err != nil {
			log.Infof("index: %s", err)
		}
	}

	log.Debugf("index: updated primary files [%s]", time.Since(start))

	return nil
}

// FlagHiddenPhotos sets the quality score of photos without valid primary file to -1.
func FlagHiddenPhotos() error {
	mutex.Index.Lock()
	defer mutex.Index.Unlock()

	start := time.Now()

	res := Db().Table("photos").
		Where("id NOT IN (SELECT photo_id FROM files WHERE file_primary = 1 AND file_missing = 0 AND file_error = '' AND deleted_at IS NULL)").
		Update("photo_quality", -1)

	switch DbDialect() {
	case MySQL:
		if res.RowsAffected > 0 {
			log.Infof("index: flagged %s as hidden or missing [%s]", english.Plural(int(res.RowsAffected), "photo", "photos"), time.Since(start))
		}
	case SQLite3:
		if res.RowsAffected > 0 {
			log.Debugf("index: flagged %s as hidden or missing [%s]", english.Plural(int(res.RowsAffected), "photo", "photos"), time.Since(start))
		}
	default:
		log.Warnf("sql: unsupported dialect %s", DbDialect())
		return nil
	}

	return res.Error
}
