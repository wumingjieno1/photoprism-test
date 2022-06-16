package entity

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/photoprism/photoprism/pkg/rnd"
)

func TestUpdate(t *testing.T) {
	var r = rand.New(rand.NewSource(time.Now().UnixNano()))
	t.Run("IDMissing", func(t *testing.T) {
		uid := rnd.GenerateUID('p')
		m := &Photo{ID: 0, PhotoUID: uid, UpdatedAt: TimeStamp(), CreatedAt: TimeStamp(), PhotoTitle: "Foo"}
		updatedAt := m.UpdatedAt

		err := Update(m, "ID", "PhotoUID")

		if err == nil {
			t.Fatal("error expected")
		}

		assert.ErrorContains(t, err, "new record")
		assert.Equal(t, m.UpdatedAt.UTC(), updatedAt.UTC())
	})
	t.Run("UIDMissing", func(t *testing.T) {
		id := 99999 + r.Intn(10000)
		m := &Photo{ID: uint(id), PhotoUID: "", UpdatedAt: TimeStamp(), CreatedAt: TimeStamp(), PhotoTitle: "Foo"}
		updatedAt := m.UpdatedAt

		err := Update(m, "ID", "PhotoUID")

		if err == nil {
			t.Fatal("error expected")
		}

		assert.ErrorContains(t, err, "record keys missing")
		assert.Equal(t, m.UpdatedAt.UTC(), updatedAt.UTC())
	})
	t.Run("NotUpdated", func(t *testing.T) {
		id := 99999 + r.Intn(10000)
		uid := rnd.GenerateUID('p')
		m := &Photo{ID: uint(id), PhotoUID: uid, UpdatedAt: time.Now(), CreatedAt: TimeStamp(), PhotoTitle: "Foo"}
		updatedAt := m.UpdatedAt

		err := Update(m, "ID", "PhotoUID")

		if err == nil {
			t.Fatal("error expected")
		}

		assert.ErrorContains(t, err, "record not found")
		assert.Greater(t, m.UpdatedAt.UTC(), updatedAt.UTC())
	})
	t.Run("Photo01", func(t *testing.T) {
		m := PhotoFixtures.Pointer("Photo01")
		updatedAt := m.UpdatedAt

		// Should be updated without any issues.
		if err := Update(m, "ID", "PhotoUID"); err != nil {
			assert.Greater(t, m.UpdatedAt.UTC(), updatedAt.UTC())
			t.Fatal(err)
			return
		} else {
			assert.Greater(t, m.UpdatedAt.UTC(), updatedAt.UTC())
			t.Logf("(1) UpdatedAt: %s -> %s", updatedAt.UTC(), m.UpdatedAt.UTC())
			t.Logf("(1) Successfully updated values")
		}

		// Tests that no error is returned on MySQL/MariaDB although
		// the number of affected rows is 0.
		if err := Update(m, "ID", "PhotoUID"); err != nil {
			assert.Greater(t, m.UpdatedAt.UTC(), updatedAt.UTC())
			t.Fatal(err)
			return
		} else {
			assert.Greater(t, m.UpdatedAt.UTC(), updatedAt.UTC())
			t.Logf("(2) UpdatedAt: %s -> %s", updatedAt.UTC(), m.UpdatedAt.UTC())
			t.Logf("(2) Successfully updated values")
		}
	})
	t.Run("NonExistentKeys", func(t *testing.T) {
		m := PhotoFixtures.Pointer("Photo01")
		m.ID = uint(99999 + r.Intn(10000))
		m.PhotoUID = rnd.GenerateUID('p')
		updatedAt := m.UpdatedAt
		if err := Update(m, "ID", "PhotoUID"); err == nil {
			t.Fatal("error expected")
			return
		} else {
			assert.ErrorContains(t, err, "record not found")
			assert.Greater(t, m.UpdatedAt.UTC(), updatedAt.UTC())
		}
	})
}
