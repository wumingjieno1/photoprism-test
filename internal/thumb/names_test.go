package thumb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName_Jpeg(t *testing.T) {
	t.Run("ResamplePng, FillCenter", func(t *testing.T) {
		assert.Equal(t, "tile_50.jpg", Tile50.Jpeg())
	})
}
