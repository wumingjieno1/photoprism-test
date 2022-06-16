package fs

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRel(t *testing.T) {
	t.Run("same", func(t *testing.T) {
		assert.Equal(t, "", RelName("/some/path", "/some/path"))
	})
	t.Run("short", func(t *testing.T) {
		assert.Equal(t, "/some/", RelName("/some/", "/some/path"))
	})
	t.Run("empty", func(t *testing.T) {
		assert.Equal(t, "", RelName("", "/some/path"))
	})
	t.Run("/some/path", func(t *testing.T) {
		assert.Equal(t, "foo/bar.baz", RelName("/some/path/foo/bar.baz", "/some/path"))
	})
	t.Run("/some/path/", func(t *testing.T) {
		assert.Equal(t, "foo/bar.baz", RelName("/some/path/foo/bar.baz", "/some/path/"))
	})
	t.Run("/some/path/bar", func(t *testing.T) {
		assert.Equal(t, "/some/path/foo/bar.baz", RelName("/some/path/foo/bar.baz", "/some/path/bar"))
	})
	t.Run("empty dir", func(t *testing.T) {
		assert.Equal(t, "/some/path/foo/bar.baz", RelName("/some/path/foo/bar.baz", ""))
	})
}

func TestFileName(t *testing.T) {
	t.Run("Test copy 3.jpg", func(t *testing.T) {
		result := FileName("testdata/Test (4).jpg", ".photoprism", Abs("testdata"), ".xmp")

		assert.Equal(t, "testdata/.photoprism/Test (4).jpg.xmp", result)
	})

	t.Run("Test (3).jpg", func(t *testing.T) {
		result := FileName("testdata/Test (4).jpg", ".photoprism", Abs("testdata"), ".xmp")

		assert.Equal(t, "testdata/.photoprism/Test (4).jpg.xmp", result)
	})

	t.Run("FOO.XMP", func(t *testing.T) {
		result := FileName("testdata/FOO.XMP", ".photoprism/sub", Abs("testdata"), ".jpeg")

		assert.Equal(t, "testdata/.photoprism/sub/FOO.XMP.jpeg", result)
	})

	t.Run("Test copy 3.jpg", func(t *testing.T) {
		tempDir := filepath.Join(os.TempDir(), HiddenPath)

		// t.Logf("TEMP DIR, ABS NAME: %s, %s", tempDir, Abs("testdata/Test (4).jpg"))

		result := FileName(Abs("testdata/Test (4).jpg"), tempDir, Abs("testdata"), ".xmp")

		assert.Equal(t, tempDir+"/Test (4).jpg.xmp", result)
	})

	t.Run("empty dir", func(t *testing.T) {
		result := FileName("testdata/FOO.XMP", "", Abs("testdata"), ".jpeg")

		assert.Equal(t, "testdata/FOO.XMP.jpeg", result)
	})
}

func TestFileNameHidden(t *testing.T) {
	t.Run("AtPrefix", func(t *testing.T) {
		assert.True(t, FileNameHidden("/some/path/@eaDir"))
	})
	t.Run("DotPrefix", func(t *testing.T) {
		assert.True(t, FileNameHidden("/some/.folder"))
	})
	t.Run("HasAts", func(t *testing.T) {
		assert.False(t, FileNameHidden("/some/path/ea@Dir"))
		assert.False(t, FileNameHidden("/some/@path/ea@Dir"))
		assert.False(t, FileNameHidden("@/eaDir"))
	})
	t.Run("HasDots", func(t *testing.T) {
		assert.False(t, FileNameHidden("/some/image.jpg."))
		assert.False(t, FileNameHidden("./.some/foo"))
	})
	t.Run("False", func(t *testing.T) {
		assert.False(t, FileNameHidden("/some/path/folder"))
	})
	t.Run("Empty", func(t *testing.T) {
		assert.False(t, FileNameHidden(""))
	})

}
