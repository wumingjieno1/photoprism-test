package fs

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirs(t *testing.T) {
	t.Run("recursive", func(t *testing.T) {
		result, err := Dirs("testdata", true, true)

		if err != nil {
			t.Fatal(err)
		}

		assert.Len(t, result, 8)
		assert.Contains(t, result, "/directory")
		assert.Contains(t, result, "/directory/subdirectory")
		assert.Contains(t, result, "/directory/subdirectory/animals")
		assert.Contains(t, result, "/linked")
	})

	t.Run("recursive no-symlinks", func(t *testing.T) {
		result, err := Dirs("testdata", true, false)

		if err != nil {
			t.Fatal(err)
		}

		assert.Contains(t, result, "/directory")
		assert.Contains(t, result, "/directory/subdirectory")
		assert.Contains(t, result, "/directory/subdirectory/animals")
		assert.Contains(t, result, "/linked")
	})

	t.Run("non-recursive", func(t *testing.T) {
		result, err := Dirs("testdata", false, true)

		if err != nil {
			t.Fatal(err)
		}

		assert.Contains(t, result, "/directory")
		assert.Contains(t, result, "/linked")
	})

	t.Run("non-recursive no-symlinks", func(t *testing.T) {
		result, err := Dirs("testdata/directory/subdirectory", false, false)

		if err != nil {
			t.Fatal(err)
		}

		assert.Len(t, result, 1)
		assert.Contains(t, result, "/animals")
	})

	t.Run("non-recursive symlinks", func(t *testing.T) {
		result, err := Dirs("testdata/linked", false, true)

		if err != nil {
			t.Fatal(err)
		}

		assert.Contains(t, result, "/photoprism")
		assert.Contains(t, result, "/self")
	})

	t.Run("no-result", func(t *testing.T) {
		result, err := Dirs("testdata/linked", false, false)

		if err != nil {
			t.Fatal(err)
		}

		assert.Empty(t, result)
	})
}

func TestFindDirs(t *testing.T) {
	t.Run("/directory", func(t *testing.T) {
		result := FindDir([]string{"/directory", "/directory/subdirectory", "/linked"})
		assert.Equal(t, "", result)
	})

	t.Run("./testdata", func(t *testing.T) {
		result := FindDir([]string{"./testdata"})
		assert.True(t, strings.HasSuffix(result, "/pkg/fs/testdata"))
	})
}
