package search

import (
	"testing"

	"github.com/photoprism/photoprism/internal/entity"
	"github.com/photoprism/photoprism/internal/form"
	"github.com/stretchr/testify/assert"
)

func TestLabels(t *testing.T) {
	t.Run("search with query", func(t *testing.T) {
		query := form.NewLabelSearch("Query:C Count:1005 Order:slug")
		result, err := Labels(query)

		if err != nil {
			t.Fatal(err)
		}

		t.Logf("results: %+v", result)

		assert.LessOrEqual(t, 2, len(result))

		for _, r := range result {
			assert.IsType(t, Label{}, r)
			assert.NotEmpty(t, r.ID)
			assert.NotEmpty(t, r.LabelName)
			assert.NotEmpty(t, r.LabelSlug)
			assert.NotEmpty(t, r.CustomSlug)

			if fix, ok := entity.LabelFixtures[r.LabelSlug]; ok {
				assert.Equal(t, fix.LabelName, r.LabelName)
				assert.Equal(t, fix.LabelSlug, r.LabelSlug)
				assert.Equal(t, fix.CustomSlug, r.CustomSlug)
			}
		}
	})

	t.Run("search for cow", func(t *testing.T) {
		query := form.NewLabelSearch("Query:cow Count:1005 Order:slug")
		result, err := Labels(query)

		if err != nil {
			t.Fatal(err)
		}

		t.Logf("results: %+v", result)

		assert.LessOrEqual(t, 1, len(result))

		for _, r := range result {
			assert.IsType(t, Label{}, r)
			assert.NotEmpty(t, r.ID)
			assert.NotEmpty(t, r.LabelName)
			assert.NotEmpty(t, r.LabelSlug)
			assert.NotEmpty(t, r.CustomSlug)

			if fix, ok := entity.LabelFixtures[r.LabelSlug]; ok {
				assert.Equal(t, fix.LabelName, r.LabelName)
				assert.Equal(t, fix.LabelSlug, r.LabelSlug)
				assert.Equal(t, fix.CustomSlug, r.CustomSlug)
			}
		}
	})
	t.Run("search for favorites", func(t *testing.T) {
		query := form.NewLabelSearch("Favorite:true Count:15")
		result, err := Labels(query)

		if err != nil {
			t.Fatal(err)
		}

		assert.LessOrEqual(t, 2, len(result))

		for _, r := range result {
			assert.True(t, r.LabelFavorite)
			assert.IsType(t, Label{}, r)
			assert.NotEmpty(t, r.ID)
			assert.NotEmpty(t, r.LabelName)
			assert.NotEmpty(t, r.LabelSlug)
			assert.NotEmpty(t, r.CustomSlug)

			if fix, ok := entity.LabelFixtures[r.LabelSlug]; ok {
				assert.Equal(t, fix.LabelName, r.LabelName)
				assert.Equal(t, fix.LabelSlug, r.LabelSlug)
				assert.Equal(t, fix.CustomSlug, r.CustomSlug)
			}
		}
	})

	t.Run("search with empty query", func(t *testing.T) {
		query := form.NewLabelSearch("")
		result, err := Labels(query)

		if err != nil {
			t.Fatal(err)
		}

		t.Log(result)
		assert.LessOrEqual(t, 3, len(result))
	})

	t.Run("search with invalid query string", func(t *testing.T) {
		query := form.NewLabelSearch("xxx:bla")
		result, err := Labels(query)

		assert.Error(t, err, "unknown filter")
		assert.Empty(t, result)
	})

	t.Run("search for ID", func(t *testing.T) {
		f := form.SearchLabels{
			Query:    "",
			UID:      "lt9k3pw1wowuy3c4",
			Slug:     "",
			Name:     "",
			All:      false,
			Favorite: false,
			Count:    0,
			Offset:   0,
			Order:    "",
		}

		result, err := Labels(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "cake", result[0].LabelSlug)
	})

	t.Run("search for label landscape", func(t *testing.T) {
		f := form.SearchLabels{
			Query: "landscape",
		}

		result, err := Labels(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "flower", result[0].LabelSlug)
	})
}
