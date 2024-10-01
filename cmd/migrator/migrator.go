package migrator

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

func RunMigrations(dbName string, driver database.Driver, migrationsPath string) {
	m, err := migrate.NewWithDatabaseInstance(
		//"file://"+migrationsPath,
		"file://"+"internal/database/migrations",
		dbName,
		driver,
	)
	if err != nil {
		panic(err)
	}

	version, dirty, err := m.Version()
	if err != nil && err != migrate.ErrNilVersion {
		log.Fatalf("Failed to get migration version: %v", err)
	}

	if dirty {
		log.Printf("Dirty database at version %d. Forcing version to %d", version, version)
		if err := m.Force(int(version)); err != nil {
			log.Fatalf("Failed to force migration version: %v", err)
		}
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migrations to apply")

			return
		}

		panic(err)
	}

	fmt.Println("migrations applied")
}
