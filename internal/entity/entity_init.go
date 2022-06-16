package entity

import (
	"os"
	"time"

	"github.com/photoprism/photoprism/pkg/clean"
)

// onReady contains init functions to be called when the
// initialization of the database is complete.
var onReady []func()

// ready executes init callbacks when the initialization of the
// database is complete.
func ready() {
	for _, init := range onReady {
		init()
	}
}

// InitDb creates database tables and inserts default fixtures as needed.
func InitDb(dropDeprecated, runFailed bool, ids []string) {
	if !HasDbProvider() {
		log.Error("migrate: no database provider")
		return
	}

	start := time.Now()

	if dropDeprecated && len(ids) == 0 {
		DeprecatedTables.Drop(Db())
	}

	Entities.Migrate(Db(), runFailed, ids)
	Entities.WaitForMigration(Db())

	CreateDefaultFixtures()

	ready()

	log.Debugf("migrate: completed in %s", time.Since(start))
}

// InitTestDb connects to and completely initializes the test database incl fixtures.
func InitTestDb(driver, dsn string) *Gorm {
	if HasDbProvider() {
		return nil
	}

	start := time.Now()

	// Set default test database driver.
	if driver == "test" || driver == "sqlite" || driver == "" || dsn == "" {
		driver = SQLite3
	}

	// Set default database DSN.
	if driver == SQLite3 {
		if dsn == "" {
			dsn = SQLiteMemoryDSN
		} else if dsn != SQLiteTestDB {
			// Continue.
		} else if err := os.Remove(dsn); err == nil {
			log.Debugf("sqlite: test file %s removed", clean.Log(dsn))
		}
	}

	log.Infof("initializing %s test db in %s", driver, dsn)

	// Create ORM instance.
	db := &Gorm{
		Driver: driver,
		Dsn:    dsn,
	}

	// Insert test fixtures.
	SetDbProvider(db)
	ResetTestFixtures()
	File{}.RegenerateIndex()

	ready()

	log.Debugf("migrate: completed in %s", time.Since(start))

	return db
}
