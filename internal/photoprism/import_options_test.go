package photoprism

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImportOptionsCopy(t *testing.T) {
	result := ImportOptionsCopy("xxx")
	assert.Equal(t, "xxx", result.Path)
	assert.Equal(t, false, result.Move)
	assert.Equal(t, false, result.RemoveDotFiles)
	assert.Equal(t, false, result.RemoveExistingFiles)
	assert.Equal(t, false, result.RemoveEmptyDirectories)
}

func TestImportOptionsMove(t *testing.T) {
	result := ImportOptionsMove("xxx")
	assert.Equal(t, "xxx", result.Path)
	assert.Equal(t, true, result.Move)
	assert.Equal(t, true, result.RemoveDotFiles)
	assert.Equal(t, true, result.RemoveExistingFiles)
	assert.Equal(t, true, result.RemoveEmptyDirectories)
}
