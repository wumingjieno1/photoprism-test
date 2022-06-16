package query

import (
	"path/filepath"

	"github.com/photoprism/photoprism/internal/entity"
	"github.com/photoprism/photoprism/internal/mutex"
	"github.com/photoprism/photoprism/pkg/fs"
)

// FoldersByPath returns a slice of folders in a given directory incl sub directories in recursive mode.
func FoldersByPath(rootName, rootPath, path string, recursive bool) (folders entity.Folders, err error) {
	dirs, err := fs.Dirs(filepath.Join(rootPath, path), recursive, true)

	if err != nil {
		return folders, err
	}

	folders = make(entity.Folders, len(dirs))

	for i, dir := range dirs {
		newFolder := entity.NewFolder(rootName, filepath.Join(path, dir), fs.BirthTime(filepath.Join(rootPath, dir)))

		if err := newFolder.Create(); err == nil {
			folders[i] = newFolder
		} else if folder := entity.FindFolder(rootName, filepath.Join(path, dir)); folder != nil {
			folders[i] = *folder
		} else {
			log.Errorf("folders: %s (create folder)", err)
		}
	}

	return folders, nil
}

// FolderCoverByUID returns a folder cover file based on the uid.
func FolderCoverByUID(uid string) (file entity.File, err error) {
	if err := Db().Where("files.file_primary = 1 AND files.file_missing = 0 AND files.file_type = 'jpg' AND files.deleted_at IS NULL").
		Joins("JOIN photos ON photos.id = files.photo_id AND photos.deleted_at IS NULL AND photos.photo_quality > -1").
		Joins("JOIN folders ON photos.photo_path = folders.path AND folders.folder_uid = ?", uid).
		Order("photos.photo_quality DESC").
		Limit(1).
		First(&file).Error; err != nil {
		return file, err
	}

	return file, nil
}

// AlbumFolders returns folders that should be added as album.
func AlbumFolders(threshold int) (folders entity.Folders, err error) {
	db := UnscopedDb().Table("folders").
		Select("folders.path, folders.root, folders.folder_uid, folders.folder_title, folders.folder_country, folders.folder_year, folders.folder_month, COUNT(photos.id) AS photo_count").
		Joins("JOIN photos ON photos.photo_path = folders.path AND photos.deleted_at IS NULL AND photos.photo_quality >= 3").
		Group("folders.path, folders.root, folders.folder_uid, folders.folder_title, folders.folder_country, folders.folder_year, folders.folder_month").
		Having("photo_count >= ?", threshold)

	if err := db.Scan(&folders).Error; err != nil {
		return folders, err
	}

	return folders, nil
}

// UpdateFolderDates updates folder year, month and day based on indexed photo metadata.
func UpdateFolderDates() error {
	mutex.Index.Lock()
	defer mutex.Index.Unlock()

	switch DbDialect() {
	case MySQL:
		return UnscopedDb().Exec(`UPDATE folders
		INNER JOIN
			(SELECT photo_path, MAX(taken_at_local) AS taken_max
			FROM photos WHERE taken_src = 'meta' AND photos.photo_quality >= 3 AND photos.deleted_at IS NULL
			GROUP BY photo_path) AS p ON folders.path = p.photo_path
		SET folders.folder_year = YEAR(taken_max), folders.folder_month = MONTH(taken_max), folders.folder_day = DAY(taken_max)
		WHERE p.taken_max IS NOT NULL`).Error
	default:
		return nil
	}
}
