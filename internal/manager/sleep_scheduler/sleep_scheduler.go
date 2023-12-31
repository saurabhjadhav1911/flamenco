package sleep_scheduler

// SPDX-License-Identifier: GPL-3.0-or-later

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"projects.blender.org/studio/flamenco/internal/manager/persistence"
	"projects.blender.org/studio/flamenco/pkg/api"
)

// Time period for checking the schedule of every worker.
const checkInterval = 1 * time.Minute

// skipWorkersInStatus has those worker statuses that should never be changed by the sleep scheduler.
var skipWorkersInStatus = map[api.WorkerStatus]bool{
	api.WorkerStatusError: true,
}

// SleepScheduler manages wake/sleep cycles of Workers.
type SleepScheduler struct {
	clock       clock.Clock
	persist     PersistenceService
	broadcaster ChangeBroadcaster
}

// New creates a new SleepScheduler.
func New(clock clock.Clock, persist PersistenceService, broadcaster ChangeBroadcaster) *SleepScheduler {
	return &SleepScheduler{
		clock:       clock,
		persist:     persist,
		broadcaster: broadcaster,
	}
}

// Run occasionally checks the sleep schedule and updates workers.
// It stops running when the context closes.
func (ss *SleepScheduler) Run(ctx context.Context) {
	log.Info().
		Str("checkInterval", checkInterval.String()).
		Msg("sleep scheduler starting")
	defer log.Info().Msg("sleep scheduler shutting down")

	waitDuration := 2 * time.Second // First check should be quickly after startup.
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(waitDuration):
			ss.CheckSchedules(ctx)
			waitDuration = checkInterval
		}
	}
}

func (ss *SleepScheduler) FetchSchedule(ctx context.Context, workerUUID string) (*persistence.SleepSchedule, error) {
	return ss.persist.FetchWorkerSleepSchedule(ctx, workerUUID)
}

// SetSleepSchedule stores the given schedule as the worker's new sleep schedule.
// The new schedule is immediately applied to the Worker.
func (ss *SleepScheduler) SetSchedule(ctx context.Context, workerUUID string, schedule *persistence.SleepSchedule) error {
	// Ensure 'start' actually preceeds 'end'.
	if schedule.StartTime.HasValue() &&
		schedule.EndTime.HasValue() &&
		schedule.EndTime.IsBefore(schedule.StartTime) {
		schedule.StartTime, schedule.EndTime = schedule.EndTime, schedule.StartTime
	}

	schedule.DaysOfWeek = cleanupDaysOfWeek(schedule.DaysOfWeek)
	schedule.NextCheck = ss.calculateNextCheck(schedule)

	if err := ss.persist.SetWorkerSleepSchedule(ctx, workerUUID, schedule); err != nil {
		return fmt.Errorf("persisting sleep schedule of worker %s: %w", workerUUID, err)
	}

	logger := addLoggerFields(zerolog.Ctx(ctx), schedule)
	logger.Info().
		Str("worker", schedule.Worker.Identifier()).
		Msg("sleep scheduler: new schedule for worker")

	return ss.ApplySleepSchedule(ctx, schedule)
}

// WorkerStatus returns the status the worker should be in right now, according to its schedule.
// If the worker has no schedule active, returns 'awake'.
func (ss *SleepScheduler) WorkerStatus(ctx context.Context, workerUUID string) (api.WorkerStatus, error) {
	schedule, err := ss.persist.FetchWorkerSleepSchedule(ctx, workerUUID)
	if err != nil {
		return "", err
	}
	return ss.scheduledWorkerStatus(schedule), nil
}

// scheduledWorkerStatus returns the expected worker status for the current date/time.
func (ss *SleepScheduler) scheduledWorkerStatus(sched *persistence.SleepSchedule) api.WorkerStatus {
	now := ss.clock.Now()
	return scheduledWorkerStatus(now, sched)
}

// Return a timestamp when the next scheck for this schedule is due.
func (ss *SleepScheduler) calculateNextCheck(schedule *persistence.SleepSchedule) time.Time {
	now := ss.clock.Now()
	return calculateNextCheck(now, schedule)
}

// ApplySleepSchedule sets worker.StatusRequested if the scheduler demands a status change.
func (ss *SleepScheduler) ApplySleepSchedule(ctx context.Context, schedule *persistence.SleepSchedule) error {
	// Find the Worker managed by this schedule.
	worker := schedule.Worker
	if worker == nil {
		err := ss.persist.FetchSleepScheduleWorker(ctx, schedule)
		if err != nil {
			return err
		}
		worker = schedule.Worker
	}

	if !ss.mayUpdateWorker(worker) {
		return nil
	}

	scheduled := ss.scheduledWorkerStatus(schedule)
	if scheduled == "" ||
		(worker.StatusRequested == scheduled && !worker.LazyStatusRequest) ||
		(worker.Status == scheduled && worker.StatusRequested == "") {
		// The worker is already in the right state, or is non-lazily requested to
		// go to the right state, so nothing else has to be done.
		return nil
	}

	logger := log.With().
		Str("worker", worker.Identifier()).
		Str("currentStatus", string(worker.Status)).
		Str("scheduledStatus", string(scheduled)).
		Logger()

	if worker.StatusRequested != "" {
		logger.Info().Str("oldStatusRequested", string(worker.StatusRequested)).
			Msg("sleep scheduler: overruling previously requested status with scheduled status")
	} else {
		logger.Info().Msg("sleep scheduler: requesting worker to switch to scheduled status")
	}

	if err := ss.updateWorkerStatus(ctx, worker, scheduled); err != nil {
		return err
	}
	return nil
}

func (ss *SleepScheduler) updateWorkerStatus(
	ctx context.Context,
	worker *persistence.Worker,
	newStatus api.WorkerStatus,
) error {
	// Sleep schedule should be adhered to immediately, no lazy requests.
	// A render task can run for hours, so better to not wait for it.
	worker.StatusChangeRequest(newStatus, false)

	err := ss.persist.SaveWorkerStatus(ctx, worker)
	if err != nil {
		return fmt.Errorf("error saving worker %s to database: %w", worker.Identifier(), err)
	}

	// Broadcast worker change via SocketIO
	ss.broadcaster.BroadcastWorkerUpdate(api.SocketIOWorkerUpdate{
		Id:     worker.UUID,
		Name:   worker.Name,
		Status: worker.Status,
		StatusChange: &api.WorkerStatusChangeRequest{
			IsLazy: false,
			Status: worker.StatusRequested,
		},
		Updated: worker.UpdatedAt,
		Version: worker.Software,
	})

	return nil
}

// CheckSchedules updates the status of all workers for which a schedule is active.
func (ss *SleepScheduler) CheckSchedules(ctx context.Context) {
	toCheck, err := ss.persist.FetchSleepSchedulesToCheck(ctx)
	if err != nil {
		log.Error().Err(err).Msg("sleep scheduler: unable to fetch sleep schedules")
		return
	}
	if len(toCheck) == 0 {
		log.Trace().Msg("sleep scheduler: no sleep schedules need checking")
		return
	}

	log.Debug().Int("numWorkers", len(toCheck)).Msg("sleep scheduler: checking worker sleep schedules")

	for _, schedule := range toCheck {
		ss.checkSchedule(ctx, schedule)
	}
}

func (ss *SleepScheduler) checkSchedule(ctx context.Context, schedule *persistence.SleepSchedule) {
	// Compute the next time to check.
	schedule.NextCheck = ss.calculateNextCheck(schedule)
	err := ss.persist.SetWorkerSleepScheduleNextCheck(ctx, schedule)
	switch {
	case errors.Is(ctx.Err(), context.Canceled):
		// Manager is shutting down, this is fine.
		return
	case err != nil:
		log.Error().
			Err(err).
			Str("worker", schedule.Worker.Identifier()).
			Msg("sleep scheduler: error refreshing worker's sleep schedule")
		return
	}

	// Apply the schedule to the worker.
	err = ss.ApplySleepSchedule(ctx, schedule)
	switch {
	case errors.Is(ctx.Err(), context.Canceled):
		// Manager is shutting down, this is fine.
	case errors.Is(err, persistence.ErrWorkerNotFound):
		// This schedule's worker cannot be found. That's fine, it could have been
		// soft-deleted (and thus foreign key constraints don't trigger deletion of
		// the sleep schedule).
		log.Debug().
			Uint("worker", schedule.WorkerID).
			Msg("sleep scheduler: sleep schedule's owning worker cannot be found; not applying the schedule")
	case err != nil:
		log.Error().
			Err(err).
			Str("worker", schedule.Worker.Identifier()).
			Msg("sleep scheduler: error applying worker's sleep schedule")
	}
}

// mayUpdateWorker determines whether the sleep scheduler is allowed to update this Worker.
func (ss *SleepScheduler) mayUpdateWorker(worker *persistence.Worker) bool {
	shouldSkip := skipWorkersInStatus[worker.Status]
	return !shouldSkip
}

func addLoggerFields(logger *zerolog.Logger, schedule *persistence.SleepSchedule) zerolog.Logger {
	logCtx := logger.With()

	if schedule.Worker != nil {
		logCtx = logCtx.Str("worker", schedule.Worker.Identifier())
	}

	logCtx = logCtx.
		Bool("isActive", schedule.IsActive).
		Str("daysOfWeek", schedule.DaysOfWeek).
		Stringer("startTime", schedule.StartTime).
		Stringer("endTime", schedule.EndTime)

	return logCtx.Logger()
}
