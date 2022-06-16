package entity

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlbum_Yaml(t *testing.T) {
	t.Run("berlin-2019", func(t *testing.T) {
		m := AlbumFixtures.Get("berlin-2019")

		if err := m.Find(); err != nil {
			t.Fatal(err)
		}

		result, err := m.Yaml()

		if err != nil {
			t.Fatal(err)
		}

		t.Logf("YAML: %s", result)
	})
	t.Run("christmas2030", func(t *testing.T) {
		m := AlbumFixtures.Get("christmas2030")

		if err := m.Find(); err != nil {
			t.Fatal(err)
		}

		result, err := m.Yaml()

		if err != nil {
			t.Fatal(err)
		}

		t.Logf("YAML: %s", result)
	})
}

func TestAlbum_SaveAsYaml(t *testing.T) {
	t.Run("berlin-2019", func(t *testing.T) {
		m := AlbumFixtures.Get("berlin-2019")

		if err := m.Find(); err != nil {
			t.Fatal(err)
		}

		fileName := m.YamlFileName("testdata")

		if err := m.SaveAsYaml(fileName); err != nil {
			t.Fatal(err)
		}

		if err := m.LoadFromYaml(fileName); err != nil {
			t.Fatal(err)
		}

		if err := os.Remove(fileName); err != nil {
			t.Fatal(err)
		}
	})
}

func TestAlbum_LoadFromYaml(t *testing.T) {
	t.Run("berlin-2020", func(t *testing.T) {
		fileName := "testdata/album/at9lxuqxpoaaaaaa.yml"

		m := Album{}

		if err := m.LoadFromYaml(fileName); err != nil {
			t.Fatal(err)
		}

		if err := m.Save(); err != nil {
			t.Fatal(err)
		}

		a := Album{AlbumUID: "at9lxuqxpoaaaaaa"}

		if err := a.Find(); err != nil {
			t.Fatal(err)
		}

		if existingYaml, err := os.ReadFile(fileName); err != nil {
			t.Fatal(err)
		} else if newYaml, err := a.Yaml(); err != nil {
			t.Fatal(err)
		} else {
			assert.Equal(t, existingYaml[:50], newYaml[:50])
			assert.Equal(t, a.AlbumUID, m.AlbumUID)
			assert.Equal(t, a.AlbumSlug, m.AlbumSlug)
			assert.Equal(t, a.AlbumType, m.AlbumType)
			assert.Equal(t, a.AlbumTitle, m.AlbumTitle)
			assert.Equal(t, a.AlbumDescription, m.AlbumDescription)
			assert.Equal(t, a.AlbumOrder, m.AlbumOrder)
			assert.Equal(t, a.AlbumCountry, m.AlbumCountry)
			assert.Equal(t, a.CreatedAt, m.CreatedAt)
			assert.Equal(t, a.UpdatedAt, m.UpdatedAt)
			assert.Equal(t, len(a.Photos), len(m.Photos))
		}
	})
}

func TestAlbum_YamlFileName(t *testing.T) {
	t.Run("berlin-2019", func(t *testing.T) {
		m := AlbumFixtures.Get("berlin-2019")

		if err := m.Find(); err != nil {
			t.Fatal(err)
		}

		fileName := m.YamlFileName("/foo/bar")

		assert.Equal(t, "/foo/bar/album/at9lxuqxpogaaba9.yml", fileName)
	})
}
