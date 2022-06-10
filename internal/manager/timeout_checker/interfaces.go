package timeout_checker

// SPDX-License-Identifier: GPL-3.0-or-later

import (
	"context"
	"time"

	"git.blender.org/flamenco/internal/manager/persistence"
	"git.blender.org/flamenco/internal/manager/task_state_machine"
	"git.blender.org/flamenco/pkg/api"
	"github.com/rs/zerolog"
)

// Generate mock implementations of these interfaces.
//go:generate go run github.com/golang/mock/mockgen -destination mocks/interfaces_mock.gen.go -package mocks git.blender.org/flamenco/internal/manager/timeout_checker PersistenceService,TaskStateMachine,LogStorage

type PersistenceService interface {
	FetchTimedOutTasks(ctx context.Context, untouchedSince time.Time) ([]*persistence.Task, error)
	FetchTimedOutWorkers(ctx context.Context, lastSeenBefore time.Time) ([]*persistence.Worker, error)
	SaveWorker(ctx context.Context, w *persistence.Worker) error
}

var _ PersistenceService = (*persistence.DB)(nil)

type TaskStateMachine interface {
	// TaskStatusChange gives a Task a new status, and handles the resulting status changes on the job.
	TaskStatusChange(ctx context.Context, task *persistence.Task, newStatus api.TaskStatus) error
	RequeueTasksOfWorker(ctx context.Context, worker *persistence.Worker, reason string) error
}

var _ TaskStateMachine = (*task_state_machine.StateMachine)(nil)

// LogStorage is used to append timeout messages to task logs.
type LogStorage interface {
	WriteTimestamped(logger zerolog.Logger, jobID, taskID string, logText string) error
}
