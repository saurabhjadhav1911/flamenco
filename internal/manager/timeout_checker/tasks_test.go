package timeout_checker

// SPDX-License-Identifier: GPL-3.0-or-later

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"git.blender.org/flamenco/internal/manager/persistence"
	"git.blender.org/flamenco/pkg/api"
)

const taskTimeout = 20 * time.Minute

func TestTimeoutCheckerTiming(t *testing.T) {
	ttc, finish, mocks := timeoutCheckerTestFixtures(t)
	defer finish()

	mocks.run(ttc)

	// Wait for the timeout checker to actually be sleeping, otherwise it could
	// have a different sleep-start time than we expect.
	time.Sleep(1 * time.Millisecond)

	// Determine the deadlines relative to the initial clock value.
	initialTime := mocks.clock.Now().UTC()
	deadlines := []time.Time{
		initialTime.Add(timeoutInitialSleep - taskTimeout),
		initialTime.Add(timeoutInitialSleep - taskTimeout + 1*timeoutCheckInterval),
		initialTime.Add(timeoutInitialSleep - taskTimeout + 2*timeoutCheckInterval),
	}

	mocks.persist.EXPECT().FetchTimedOutWorkers(mocks.ctx, gomock.Any()).AnyTimes().Return(nil, nil)

	// Expect three fetches, one after the initial sleep time, and two a regular interval later.
	fetchTimes := make([]time.Time, len(deadlines))
	firstCall := mocks.persist.EXPECT().FetchTimedOutTasks(mocks.ctx, deadlines[0]).
		DoAndReturn(func(ctx context.Context, timeout time.Time) ([]*persistence.Task, error) {
			fetchTimes[0] = mocks.clock.Now().UTC()
			return []*persistence.Task{}, nil
		})

	secondCall := mocks.persist.EXPECT().FetchTimedOutTasks(mocks.ctx, deadlines[1]).
		DoAndReturn(func(ctx context.Context, timeout time.Time) ([]*persistence.Task, error) {
			fetchTimes[1] = mocks.clock.Now().UTC()
			// Return a database error. This shouldn't break the check loop.
			return []*persistence.Task{}, errors.New("testing what errors do")
		}).
		After(firstCall)

	mocks.persist.EXPECT().FetchTimedOutTasks(mocks.ctx, deadlines[2]).
		DoAndReturn(func(ctx context.Context, timeout time.Time) ([]*persistence.Task, error) {
			fetchTimes[2] = mocks.clock.Now().UTC()
			return []*persistence.Task{}, nil
		}).
		After(secondCall)

	mocks.clock.Add(2 * time.Minute) // Should still be sleeping.
	mocks.clock.Add(2 * time.Minute) // Should still be sleeping.
	mocks.clock.Add(time.Minute)     // Should trigger the first fetch.
	mocks.clock.Add(time.Minute)     // Should trigger the second fetch.
	mocks.clock.Add(time.Minute)     // Should trigger the third fetch.

	// Wait for the timeout checker to actually run & hit the expected calls.
	time.Sleep(1 * time.Millisecond)

	for idx, fetchTime := range fetchTimes {
		// Check for zero values first, because they can be a bit confusing in the assert.Equal() logs.
		if !assert.Falsef(t, fetchTime.IsZero(), "fetchTime[%d] should not be zero", idx) {
			continue
		}
		expect := initialTime.Add(timeoutInitialSleep + time.Duration(idx)*timeoutCheckInterval)
		assert.Equalf(t, expect, fetchTime, "fetchTime[%d] not as expected", idx)
	}
}

func TestTaskTimeout(t *testing.T) {
	// Canary test: if these constants do not have the expected value, the test
	// will fail rather cryptically.
	if !assert.Equal(t, 5*time.Minute, timeoutInitialSleep, "timeoutInitialSleep does not have the expected value") ||
		!assert.Equal(t, 1*time.Minute, timeoutCheckInterval, "timeoutCheckInterval does not have the expected value") {
		t.FailNow()
	}

	ttc, finish, mocks := timeoutCheckerTestFixtures(t)
	defer finish()

	mocks.run(ttc)

	// Wait for the timeout checker to actually be sleeping, otherwise it could
	// have a different sleep-start time than we expect.
	time.Sleep(1 * time.Millisecond)

	lastTime := mocks.clock.Now().UTC().Add(-1 * time.Hour)

	job := persistence.Job{UUID: "JOB-UUID"}
	worker := persistence.Worker{
		UUID:  "WORKER-UUID",
		Name:  "Tester",
		Model: gorm.Model{ID: 47},
	}
	taskUnassigned := persistence.Task{
		UUID:          "TASK-UUID-UNASSIGNED",
		Job:           &job,
		LastTouchedAt: lastTime,
	}
	taskUnknownWorker := persistence.Task{
		UUID:          "TASK-UUID-UNKNOWN",
		Job:           &job,
		LastTouchedAt: lastTime,
		WorkerID:      &worker.ID,
	}
	taskAssigned := persistence.Task{
		UUID:          "TASK-UUID-ASSIGNED",
		Job:           &job,
		LastTouchedAt: lastTime,
		WorkerID:      &worker.ID,
		Worker:        &worker,
	}

	mocks.persist.EXPECT().FetchTimedOutWorkers(mocks.ctx, gomock.Any()).AnyTimes().Return(nil, nil)

	mocks.persist.EXPECT().FetchTimedOutTasks(mocks.ctx, gomock.Any()).
		Return([]*persistence.Task{&taskUnassigned, &taskUnknownWorker, &taskAssigned}, nil)

	mocks.taskStateMachine.EXPECT().TaskStatusChange(mocks.ctx, &taskUnassigned, api.TaskStatusFailed)
	mocks.taskStateMachine.EXPECT().TaskStatusChange(mocks.ctx, &taskUnknownWorker, api.TaskStatusFailed)
	mocks.taskStateMachine.EXPECT().TaskStatusChange(mocks.ctx, &taskAssigned, api.TaskStatusFailed)

	mocks.logStorage.EXPECT().WriteTimestamped(gomock.Any(), job.UUID, taskUnassigned.UUID,
		"Task timed out. It was assigned to worker -unassigned-, but untouched since 1969-12-31T23:00:00Z")
	mocks.logStorage.EXPECT().WriteTimestamped(gomock.Any(), job.UUID, taskUnknownWorker.UUID,
		"Task timed out. It was assigned to worker -unknown-, but untouched since 1969-12-31T23:00:00Z")
	mocks.logStorage.EXPECT().WriteTimestamped(gomock.Any(), job.UUID, taskAssigned.UUID,
		"Task timed out. It was assigned to worker Tester (WORKER-UUID), but untouched since 1969-12-31T23:00:00Z")

	// All the timeouts should be handled after the initial sleep.
	mocks.clock.Add(timeoutInitialSleep)
}
