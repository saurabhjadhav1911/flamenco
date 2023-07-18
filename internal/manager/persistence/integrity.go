package persistence

// SPDX-License-Identifier: GPL-3.0-or-later

import (
	"context"
	"errors"
	"time"

	"github.com/rs/zerolog/log"
)

var ErrIntegrity = errors.New("database integrity check failed")

const (
	integrityCheckTimeout = 2 * time.Second
	integrityCheckPeriod  = 1 * time.Hour
)

type PragmaIntegrityCheckResult struct {
	Description string `gorm:"column:integrity_check"`
}

type PragmaForeignKeyCheckResult struct {
	Table  string `gorm:"column:table"`
	RowID  int    `gorm:"column:rowid"`
	Parent string `gorm:"column:parent"`
	FKID   int    `gorm:"column:fkid"`
}

// PeriodicIntegrityCheck periodically checks the database integrity.
// This function only returns when the context is done.
func (db *DB) PeriodicIntegrityCheck(ctx context.Context, onErrorCallback func()) {
	log.Debug().Msg("database: periodic integrity check loop starting")
	defer log.Debug().Msg("database: periodic integrity check loop stopping")

	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(integrityCheckPeriod):
		}

		ok := db.performIntegrityCheck(ctx)
		if !ok {
			log.Error().Msg("database: periodic integrity check failed")
			onErrorCallback()
		}
	}
}

// performIntegrityCheck uses a few 'pragma' SQL statements to do some integrity checking.
// Returns true on OK, false if there was an issue. Issues are always logged.
func (db *DB) performIntegrityCheck(ctx context.Context) (ok bool) {
	checkCtx, cancel := context.WithTimeout(ctx, integrityCheckTimeout)
	defer cancel()

	if !db.pragmaIntegrityCheck(checkCtx) {
		return false
	}
	return db.pragmaForeignKeyCheck(checkCtx)
}

// pragmaIntegrityCheck checks database file integrity. This does not include
// foreign key checks.
//
// Returns true on OK, false if there was an issue. Issues are always logged.
//
// See https: //www.sqlite.org/pragma.html#pragma_integrity_check
func (db *DB) pragmaIntegrityCheck(ctx context.Context) (ok bool) {
	var issues []PragmaIntegrityCheckResult

	tx := db.gormDB.WithContext(ctx).
		Raw("PRAGMA integrity_check").
		Scan(&issues)
	if tx.Error != nil {
		log.Error().Err(tx.Error).Msg("database error checking integrity")
		return false
	}

	switch len(issues) {
	case 0:
		log.Warn().Msg("database integrity check returned nothing, expected explicit 'ok'; treating as an implicit 'ok'")
		return true
	case 1:
		if issues[0].Description == "ok" {
			log.Debug().Msg("database integrity check ok")
			return true
		}
	}

	log.Error().Int("num_issues", len(issues)).Msg("database integrity check failed")
	for _, issue := range issues {
		log.Error().
			Str("description", issue.Description).
			Msg("database integrity check failure")
	}

	return false
}

// pragmaForeignKeyCheck checks whether all foreign key constraints are still valid.
//
// SQLite has optional foreign key relations, so even though Flamenco Manager
// always enables these on startup, at some point there could be some issue
// causing these checks to be skipped.
//
// Returns true on OK, false if there was an issue. Issues are always logged.
//
// See https: //www.sqlite.org/pragma.html#pragma_foreign_key_check
func (db *DB) pragmaForeignKeyCheck(ctx context.Context) (ok bool) {
	var issues []PragmaForeignKeyCheckResult

	tx := db.gormDB.WithContext(ctx).
		Raw("PRAGMA foreign_key_check").
		Scan(&issues)
	if tx.Error != nil {
		log.Error().Err(tx.Error).Msg("database error checking foreign keys")
		return false
	}

	if len(issues) == 0 {
		log.Debug().Msg("database foreign key check ok")
		return true
	}

	log.Error().Int("num_issues", len(issues)).Msg("database foreign key check failed")
	for _, issue := range issues {
		log.Error().
			Str("table", issue.Table).
			Int("rowid", issue.RowID).
			Str("parent", issue.Parent).
			Int("fkid", issue.FKID).
			Msg("database foreign key relation missing")
	}

	return false
}
