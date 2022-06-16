package meta

import (
	"fmt"
	"path/filepath"
	"runtime/debug"

	"github.com/photoprism/photoprism/pkg/clean"
)

// XMP parses an XMP file and returns a Data struct.
func XMP(fileName string) (data Data, err error) {
	err = data.XMP(fileName)

	return data, err
}

// XMP parses an XMP file and returns a Data struct.
func (data *Data) XMP(fileName string) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("metadata: %s in %s (xmp panic)\nstack: %s", e, clean.Log(filepath.Base(fileName)), debug.Stack())
		}
	}()

	doc := XmpDocument{}

	if err := doc.Load(fileName); err != nil {
		return fmt.Errorf("metadata: cannot read %s (xmp)", clean.Log(filepath.Base(fileName)))
	}

	if doc.Title() != "" {
		data.Title = doc.Title()
	}

	if doc.Artist() != "" {
		data.Artist = doc.Artist()
	}

	if doc.Description() != "" {
		data.Description = doc.Description()
	}

	if doc.Copyright() != "" {
		data.Copyright = doc.Copyright()
	}

	if doc.CameraMake() != "" {
		data.CameraMake = doc.CameraMake()
	}

	if doc.CameraModel() != "" {
		data.CameraModel = doc.CameraModel()
	}

	if doc.LensModel() != "" {
		data.LensModel = doc.LensModel()
	}

	if takenAt := doc.TakenAt(); !takenAt.IsZero() {
		data.TakenAt = takenAt
	}

	if len(doc.Keywords()) != 0 {
		data.AddKeywords(doc.Keywords())
	}

	return nil
}
