package entity

import (
	"os"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
)

func TestMySQL8(t *testing.T) {
	dbDsn := os.Getenv("PHOTOPRISM_TEST_DSN_MYSQL8")

	if dbDsn == "" {
		t.Skip("skipping MySQL 8 test: PHOTOPRISM_TEST_DSN_MYSQL8 is not set")
	}

	dbDriver := MySQL
	db, err := gorm.Open(dbDriver, dbDsn)

	if err != nil || db == nil {
		for i := 1; i <= 5; i++ {
			db, err = gorm.Open(dbDriver, dbDsn)

			if db != nil && err == nil {
				break
			}

			time.Sleep(5 * time.Second)
		}

		if err != nil || db == nil {
			t.Fatal(err)
		}
	}

	defer db.Close()

	db.LogMode(false)

	DeprecatedTables.Drop(db)
	Entities.Drop(db)

	// First migration.
	Entities.Migrate(db, false, nil)
	Entities.WaitForMigration(db)

	// Second migration.
	Entities.Migrate(db, false, nil)
	Entities.WaitForMigration(db)

	// Third migration with force flag.
	Entities.Migrate(db, true, []string{"20211121-094727"})
	Entities.WaitForMigration(db)
}
