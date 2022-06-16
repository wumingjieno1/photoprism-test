package entity

import (
	"testing"
)

func TestDeleteOrphanPeople(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		if count, err := DeleteOrphanPeople(); err != nil {
			t.Fatal(err)
		} else {
			t.Logf("deleted %d faces", count)
		}
	})
}
