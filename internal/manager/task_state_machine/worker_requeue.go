package task_state_machine

// SPDX-License-Identifier: GPL-3.0-or-later

import (
	"context"

	"github.com/rs/zerolog/log"
	"projects.blender.org/studio/flamenco/internal/manager/persistence"
	"projects.blender.org/studio/flamenco/pkg/api"
)

// RequeueActiveTasksOfWorker re-queues all active tasks (should be max one) of this worker.
//
// `reason`: a string that can be appended to text like "Task requeued because "
func (sm *StateMachine) RequeueActiveTasksOfWorker(
	ctx context.Context,
	worker *persistence.Worker,
	reason string,
) error {
	// Fetch the tasks to update.
	tasks, err := sm.persist.FetchTasksOfWorkerInStatus(
		ctx, worker, api.TaskStatusActive)
	if err != nil {
		return err
	}

	return sm.requeueTasksOfWorker(ctx, tasks, worker, reason)
}

// RequeueFailedTasksOfWorkerOfJob re-queues all failed tasks of this worker on this job.
//
// `reason`: a string that can be appended to text like "Task requeued because "
func (sm *StateMachine) RequeueFailedTasksOfWorkerOfJob(
	ctx context.Context,
	worker *persistence.Worker,
	job *persistence.Job,
	reason string,
) error {
	// Fetch the tasks to update.
	tasks, err := sm.persist.FetchTasksOfWorkerInStatusOfJob(
		ctx, worker, api.TaskStatusFailed, job)
	if err != nil {
		return err
	}

	return sm.requeueTasksOfWorker(ctx, tasks, worker, reason)
}

func (sm *StateMachine) requeueTasksOfWorker(
	ctx context.Context,
	tasks []*persistence.Task,
	worker *persistence.Worker,
	reason string,
) error {
	logger := log.With().
		Str("worker", worker.UUID).
		Str("reason", reason).
		Logger()

	// Run each task change through the task state machine.
	var lastErr error
	for _, task := range tasks {
		logger.Info().
			Str("task", task.UUID).
			Msg("re-queueing task")

			// Write to task activity that it got requeued because of worker sign-off.
		task.Activity = "Task was requeued by Manager because " + reason
		if err := sm.persist.SaveTaskActivity(ctx, task); err != nil {
			logger.Warn().Err(err).
				Str("task", task.UUID).
				Str("reason", reason).
				Str("activity", task.Activity).
				Msg("error saving task activity to database")
			lastErr = err
		}

		if err := sm.TaskStatusChange(ctx, task, api.TaskStatusQueued); err != nil {
			logger.Warn().Err(err).
				Str("task", task.UUID).
				Str("reason", reason).
				Msg("error queueing task")
			lastErr = err
		}

		// The error is already logged by the log storage.
		_ = sm.logStorage.WriteTimestamped(logger, task.Job.UUID, task.UUID, task.Activity)
	}

	return lastErr
}
