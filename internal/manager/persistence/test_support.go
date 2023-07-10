// Package persistence provides the database interface for Flamenco Manager.
package persistence

// SPDX-License-Identifier: GPL-3.0-or-later

import (
	"context"
	"database/sql"
	"os"
	"testing"
	"time"

	"git.blender.org/flamenco/internal/uuid"
	"git.blender.org/flamenco/pkg/api"
	"github.com/glebarez/sqlite"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

// Change this to a filename if you want to run a single test and inspect the
// resulting database.
const TestDSN = "file::memory:"

func CreateTestDB(t *testing.T) (db *DB, closer func()) {
	// Delete the SQLite file if it exists on disk.
	if _, err := os.Stat(TestDSN); err == nil {
		if err := os.Remove(TestDSN); err != nil {
			t.Fatalf("unable to remove %s: %v", TestDSN, err)
		}
	}

	var err error

	dblogger := NewDBLogger(log.Level(zerolog.TraceLevel).Output(os.Stdout))

	// Open the database ourselves, so that we have a low-level connection that
	// can be closed when the unit test is done running.
	sqliteConn, err := sql.Open(sqlite.DriverName, TestDSN)
	if err != nil {
		t.Fatalf("opening SQLite connection: %v", err)
	}

	config := gorm.Config{
		Logger:   dblogger,
		ConnPool: sqliteConn,
		NowFunc:  nowFunc,
	}

	db, err = openDBWithConfig(TestDSN, &config)
	if err != nil {
		t.Fatalf("opening DB: %v", err)
	}

	err = db.migrate()
	if err != nil {
		t.Fatalf("migrating DB: %v", err)
	}

	closer = func() {
		if err := sqliteConn.Close(); err != nil {
			t.Fatalf("closing DB: %v", err)
		}
	}

	return db, closer
}

// persistenceTestFixtures creates a test database and returns it and a context.
// Tests should call the returned cancel function when they're done.
func persistenceTestFixtures(t *testing.T, testContextTimeout time.Duration) (context.Context, context.CancelFunc, *DB) {
	db, dbCloser := CreateTestDB(t)

	var (
		ctx       context.Context
		ctxCancel context.CancelFunc
	)
	if testContextTimeout > 0 {
		ctx, ctxCancel = context.WithTimeout(context.Background(), testContextTimeout)
	} else {
		ctx = context.Background()
		ctxCancel = func() {}
	}

	cancel := func() {
		ctxCancel()
		dbCloser()
	}

	return ctx, cancel, db
}

type WorkerTestFixture struct {
	db   *DB
	ctx  context.Context
	done func()

	worker *Worker
	tag    *WorkerTag
}

func workerTestFixtures(t *testing.T, testContextTimeout time.Duration) WorkerTestFixture {
	ctx, cancel, db := persistenceTestFixtures(t, testContextTimeout)

	w := Worker{
		UUID:               uuid.New(),
		Name:               "дрон",
		Address:            "fe80::5054:ff:fede:2ad7",
		Platform:           "linux",
		Software:           "3.0",
		Status:             api.WorkerStatusAwake,
		SupportedTaskTypes: "blender,ffmpeg,file-management",
	}

	wc := WorkerTag{
		UUID:        uuid.New(),
		Name:        "arbejdsklynge",
		Description: "Worker tag in Danish",
	}

	require.NoError(t, db.CreateWorker(ctx, &w))
	require.NoError(t, db.CreateWorkerTag(ctx, &wc))

	return WorkerTestFixture{
		db:   db,
		ctx:  ctx,
		done: cancel,

		worker: &w,
		tag:    &wc,
	}
}
