package timeout_checker

// SPDX-License-Identifier: GPL-3.0-or-later

import (
	"testing"
	"time"

	"git.blender.org/flamenco/internal/manager/persistence"
	"git.blender.org/flamenco/pkg/api"
	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
)

const workerTimeout = 20 * time.Minute

func TestWorkerTimeout(t *testing.T) {
	ttc, finish, mocks := timeoutCheckerTestFixtures(t)
	defer finish()

	mocks.run(ttc)

	// Wait for the timeout checker to actually be sleeping, otherwise it could
	// have a different sleep-start time than we expect.
	time.Sleep(1 * time.Millisecond)

	lastSeenAt := mocks.clock.Now().UTC().Add(-1 * time.Hour)

	worker := persistence.Worker{
		UUID:            "WORKER-UUID",
		Name:            "Tester",
		Model:           gorm.Model{ID: 47},
		LastSeenAt:      lastSeenAt,
		Status:          api.WorkerStatusAsleep,
		StatusRequested: api.WorkerStatusAwake,
	}

	mocks.persist.EXPECT().FetchTimedOutTasks(mocks.ctx, gomock.Any()).Return([]*persistence.Task{}, nil)
	mocks.persist.EXPECT().FetchTimedOutWorkers(mocks.ctx, gomock.Any()).
		Return([]*persistence.Worker{&worker}, nil)

	persistedWorker := worker
	persistedWorker.Status = api.WorkerStatusError
	// Any queued up status change should be cleared, as the Worker is not allowed
	// to change into anything until this timeout has been address.
	persistedWorker.StatusChangeClear()
	mocks.persist.EXPECT().SaveWorker(mocks.ctx, &persistedWorker).Return(nil)

	// TODO: expect all tasks assigned to the worker to get requeued.

	// All the timeouts should be handled after the initial sleep.
	mocks.clock.Add(timeoutInitialSleep)
}
