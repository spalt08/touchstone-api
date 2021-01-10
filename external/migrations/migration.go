package migrations

import (
	"touchstone-api/pkg/utils/logger"

	"github.com/go-pg/migrations/v8"
	"github.com/go-pg/pg/v10"
)

// RunMigrations will apply all available migrations
func RunMigrations(db *pg.DB) {
	version, err := migrations.Version(db)

	if err != nil {
		_, _, err := migrations.Run(db, "init")
		logger.Info("[MIGRATION] Migration table created")

		if err != nil {
			panic(err)
		}
	}

	logger.Info("[MIGRATION] Current version: ", version)
	_, newVersion, err := migrations.Run(db, "up")

	if err != nil {
		logger.Info("An error happend during migration")
		panic(err)
	}

	if newVersion != version {
		logger.Info("[MIGRATION] Migrated from version %d to %d\n", version, newVersion)
	} else {
		logger.Info("[MIGRATION] Up to date")
	}
}
