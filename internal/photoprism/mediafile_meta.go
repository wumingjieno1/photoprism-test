package photoprism

import (
	"fmt"
	"path/filepath"

	"github.com/photoprism/photoprism/internal/entity"
	"github.com/photoprism/photoprism/internal/meta"

	"github.com/photoprism/photoprism/pkg/clean"
	"github.com/photoprism/photoprism/pkg/fs"
)

// HasSidecarJson returns true if this file has or is a json sidecar file.
func (m *MediaFile) HasSidecarJson() bool {
	if m.IsJson() {
		return true
	}

	return fs.JsonFile.FindFirst(m.FileName(), []string{Config().SidecarPath(), fs.HiddenPath}, Config().OriginalsPath(), false) != ""
}

// SidecarJsonName returns the corresponding JSON sidecar file name as used by Google Photos (and potentially other apps).
func (m *MediaFile) SidecarJsonName() string {
	jsonName := m.fileName + ".json"

	if fs.FileExistsNotEmpty(jsonName) {
		return jsonName
	}

	return ""
}

// ExifToolJsonName returns the cached ExifTool metadata file name.
func (m *MediaFile) ExifToolJsonName() (string, error) {
	if Config().DisableExifTool() {
		return "", fmt.Errorf("media: exiftool json files disabled")
	}

	return CacheName(m.Hash(), "json", "exiftool.json")
}

// NeedsExifToolJson tests if an ExifTool JSON file needs to be created.
func (m *MediaFile) NeedsExifToolJson() bool {
	if m.Root() == entity.RootSidecar || !m.IsMedia() {
		return false
	}

	jsonName, err := m.ExifToolJsonName()

	if err != nil {
		return false
	}

	return !fs.FileExists(jsonName)
}

// ReadExifToolJson reads metadata from a cached ExifTool JSON file.
func (m *MediaFile) ReadExifToolJson() error {
	jsonName, err := m.ExifToolJsonName()

	if err != nil {
		return err
	}

	return m.metaData.JSON(jsonName, "")
}

// MetaData returns exif meta data of a media file.
func (m *MediaFile) MetaData() (result meta.Data) {
	m.metaOnce.Do(func() {
		var err error

		if m.ExifSupported() {
			err = m.metaData.Exif(m.FileName(), m.FileType(), Config().ExifBruteForce())
		} else {
			err = fmt.Errorf("exif not supported")
		}

		// Parse regular JSON sidecar files ("img_1234.json")
		if !m.IsSidecar() {
			if jsonFiles := fs.JsonFile.FindAll(m.FileName(), []string{Config().SidecarPath(), fs.HiddenPath}, Config().OriginalsPath(), false); len(jsonFiles) == 0 {
				log.Tracef("metadata: found no additional sidecar file for %s", clean.Log(filepath.Base(m.FileName())))
			} else {
				for _, jsonFile := range jsonFiles {
					jsonErr := m.metaData.JSON(jsonFile, m.BaseName())

					if jsonErr != nil {
						log.Debug(jsonErr)
					} else {
						err = nil
					}
				}
			}

			if jsonErr := m.ReadExifToolJson(); jsonErr != nil {
				log.Debug(jsonErr)
			} else {
				err = nil
			}
		}

		if err != nil {
			m.metaData.Error = err
			log.Debugf("metadata: %s in %s", err, clean.Log(m.BaseName()))
		}
	})

	return m.metaData
}
