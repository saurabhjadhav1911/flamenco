// Package persistence provides the database interface for Flamenco Manager.
package persistence

// SPDX-License-Identifier: GPL-3.0-or-later

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"git.blender.org/flamenco/internal/uuid"
	"git.blender.org/flamenco/pkg/api"
)

func TestCreateFetchWorker(t *testing.T) {
	ctx, cancel, db := persistenceTestFixtures(t, 1*time.Second)
	defer cancel()

	// Test fetching non-existent worker
	fetchedWorker, err := db.FetchWorker(ctx, "dabf67a1-b591-4232-bf73-0b8de2a9488e")
	assert.ErrorIs(t, err, ErrWorkerNotFound)
	assert.Nil(t, fetchedWorker)

	// Test existing worker
	w := Worker{
		UUID:               uuid.New(),
		Name:               "дрон",
		Address:            "fe80::5054:ff:fede:2ad7",
		Platform:           "linux",
		Software:           "3.0",
		Status:             api.WorkerStatusAwake,
		SupportedTaskTypes: "blender,ffmpeg,file-management",
	}

	err = db.CreateWorker(ctx, &w)
	assert.NoError(t, err)

	fetchedWorker, err = db.FetchWorker(ctx, w.UUID)
	assert.NoError(t, err)
	assert.NotNil(t, fetchedWorker)

	// Test contents of fetched job
	assert.Equal(t, w.UUID, fetchedWorker.UUID)
	assert.Equal(t, w.Name, fetchedWorker.Name)
	assert.Equal(t, w.Address, fetchedWorker.Address)
	assert.Equal(t, w.Platform, fetchedWorker.Platform)
	assert.Equal(t, w.Software, fetchedWorker.Software)
	assert.Equal(t, w.Status, fetchedWorker.Status)

	assert.EqualValues(t, w.SupportedTaskTypes, fetchedWorker.SupportedTaskTypes)
}

func TestSaveWorker(t *testing.T) {
	ctx, cancel, db := persistenceTestFixtures(t, 1*time.Second)
	defer cancel()

	w := Worker{
		UUID:               uuid.New(),
		Name:               "дрон",
		Address:            "fe80::5054:ff:fede:2ad7",
		Platform:           "linux",
		Software:           "3.0",
		Status:             api.WorkerStatusAwake,
		SupportedTaskTypes: "blender,ffmpeg,file-management",
	}

	err := db.CreateWorker(ctx, &w)
	assert.NoError(t, err)

	fetchedWorker, err := db.FetchWorker(ctx, w.UUID)
	assert.NoError(t, err)
	assert.NotNil(t, fetchedWorker)

	// Update all updatable fields of the Worker
	updatedWorker := *fetchedWorker
	updatedWorker.Name = "7 မှ 9"
	updatedWorker.Address = "fe80::cafe:f00d"
	updatedWorker.Platform = "windows"
	updatedWorker.Software = "3.1"
	updatedWorker.Status = api.WorkerStatusAsleep
	updatedWorker.SupportedTaskTypes = "blender,ffmpeg,file-management,misc"

	// Saving only the status should just do that.
	err = db.SaveWorkerStatus(ctx, &updatedWorker)
	assert.NoError(t, err)
	assert.Equal(t, "7 မှ 9", updatedWorker.Name, "Saving status should not touch the name")

	// Check saved worker
	fetchedWorker, err = db.FetchWorker(ctx, w.UUID)
	assert.NoError(t, err)
	assert.NotNil(t, fetchedWorker)
	assert.Equal(t, updatedWorker.Status, fetchedWorker.Status, "new status should have been saved")
	assert.NotEqual(t, updatedWorker.Name, fetchedWorker.Name, "non-status fields should not have been updated")

	// Saving the entire worker should save everything.
	err = db.SaveWorker(ctx, &updatedWorker)
	assert.NoError(t, err)

	// Check saved worker
	fetchedWorker, err = db.FetchWorker(ctx, w.UUID)
	assert.NoError(t, err)
	assert.NotNil(t, fetchedWorker)
	assert.Equal(t, updatedWorker.Status, fetchedWorker.Status, "new status should have been saved")
	assert.Equal(t, updatedWorker.Name, fetchedWorker.Name, "non-status fields should also have been updated")
	assert.Equal(t, updatedWorker.Software, fetchedWorker.Software, "non-status fields should also have been updated")
}

func TestFetchWorkers(t *testing.T) {
	ctx, cancel, db := persistenceTestFixtures(t, 1*time.Second)
	defer cancel()

	// No workers
	workers, err := db.FetchWorkers(ctx)
	if !assert.NoError(t, err) {
		t.Fatal("error fetching empty list of workers, no use in continuing the test")
	}
	assert.Empty(t, workers)

	linuxWorker := Worker{
		UUID:               uuid.New(),
		Name:               "дрон",
		Address:            "fe80::5054:ff:fede:2ad7",
		Platform:           "linux",
		Software:           "3.0",
		Status:             api.WorkerStatusAwake,
		SupportedTaskTypes: "blender,ffmpeg,file-management",
	}

	// One worker:
	err = db.CreateWorker(ctx, &linuxWorker)
	assert.NoError(t, err)
	assert.Equal(t, time.Now().UTC().Location(), linuxWorker.CreatedAt.Location(),
		"Timestamps should be using UTC timezone")

	workers, err = db.FetchWorkers(ctx)
	assert.NoError(t, err)
	if assert.Len(t, workers, 1) {
		// FIXME: this fails, because the fetched timestamps have nil location instead of UTC.
		// assert.Equal(t, time.Now().UTC().Location(), workers[0].CreatedAt.Location(),
		// 	"Timestamps should be using UTC timezone")

		assert.Equal(t, linuxWorker.UUID, workers[0].UUID)
		assert.Equal(t, linuxWorker.Name, workers[0].Name)
		assert.Equal(t, linuxWorker.Address, workers[0].Address)
		assert.Equal(t, linuxWorker.Status, workers[0].Status)
		assert.Equal(t, linuxWorker.SupportedTaskTypes, workers[0].SupportedTaskTypes)
	}

	// Two workers:
	windowsWorker := Worker{
		UUID:               uuid.New(),
		Name:               "очиститель окон",
		Address:            "fe80::c000:d000:::3",
		Platform:           "windows",
		Software:           "3.0",
		Status:             api.WorkerStatusOffline,
		SupportedTaskTypes: "blender,ffmpeg,file-management",
	}
	err = db.CreateWorker(ctx, &windowsWorker)
	assert.NoError(t, err)

	workers, err = db.FetchWorkers(ctx)
	assert.NoError(t, err)
	if assert.Len(t, workers, 2) {
		assert.Equal(t, linuxWorker.UUID, workers[0].UUID)
		assert.Equal(t, windowsWorker.UUID, workers[1].UUID)
	}
}
