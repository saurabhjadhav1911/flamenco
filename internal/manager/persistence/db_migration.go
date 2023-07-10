package persistence

// SPDX-License-Identifier: GPL-3.0-or-later

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

func (db *DB) migrate() error {
	log.Debug().Msg("auto-migrating database")

	// There is an issue with the GORM auto-migration, in that it doesn't always
	// disable foreign key constraints when it should. Due to limitations of
	// SQLite, not all 'alter table' commands you'd want to use are available. As
	// a workaround, these steps are performed:
	//
	// 1. create a new table with the desired schema,
	// 2. copy the data over,
	// 3. drop the old table,
	// 4. rename the new table to the old name.
	//
	// Step #3 will wreak havoc with the database when foreign key constraint
	// checks are active.

	if err := db.pragmaForeignKeys(false); err != nil {
		return fmt.Errorf("disabling foreign key checks before auto-migration: %w", err)
	}
	defer func() {
		err := db.pragmaForeignKeys(true)
		if err != nil {
			// There is no way that Flamenco Manager should be runnign with foreign key checks disabled.
			log.Fatal().Err(err).Msg("re-enabling foreign key checks after auto-migration failed")
		}
	}()

	err := db.gormDB.AutoMigrate(
		&Job{},
		&JobBlock{},
		&JobStorageInfo{},
		&LastRendered{},
		&SleepSchedule{},
		&Task{},
		&TaskFailure{},
		&Worker{},
		&WorkerTag{},
	)
	if err != nil {
		return fmt.Errorf("failed to automigrate database: %v", err)
	}
	return nil
}
