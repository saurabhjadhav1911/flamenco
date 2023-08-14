package api_impl

// SPDX-License-Identifier: GPL-3.0-or-later

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"projects.blender.org/studio/flamenco/internal/manager/persistence"
	"projects.blender.org/studio/flamenco/internal/manager/webupdates"
	"projects.blender.org/studio/flamenco/internal/uuid"
	"projects.blender.org/studio/flamenco/pkg/api"
)

func (f *Flamenco) FetchWorkers(e echo.Context) error {
	logger := requestLogger(e)
	dbWorkers, err := f.persist.FetchWorkers(e.Request().Context())
	if err != nil {
		logger.Error().Err(err).Msg("fetching all workers")
		return sendAPIError(e, http.StatusInternalServerError, "error fetching workers: %v", err)
	}

	apiWorkers := make([]api.WorkerSummary, len(dbWorkers))
	for i := range dbWorkers {
		apiWorkers[i] = workerSummary(*dbWorkers[i])
	}

	logger.Debug().Msg("fetched all workers")
	return e.JSON(http.StatusOK, api.WorkerList{
		Workers: apiWorkers,
	})
}

func (f *Flamenco) FetchWorker(e echo.Context, workerUUID string) error {
	logger := requestLogger(e)
	logger = logger.With().Str("worker", workerUUID).Logger()

	if !uuid.IsValid(workerUUID) {
		return sendAPIError(e, http.StatusBadRequest, "not a valid UUID")
	}

	ctx := e.Request().Context()
	dbWorker, err := f.persist.FetchWorker(ctx, workerUUID)
	if errors.Is(err, persistence.ErrWorkerNotFound) {
		logger.Debug().Msg("non-existent worker requested")
		return sendAPIError(e, http.StatusNotFound, "worker %q not found", workerUUID)
	}
	if err != nil {
		logger.Error().Err(err).Msg("fetching worker")
		return sendAPIError(e, http.StatusInternalServerError, "error fetching worker: %v", err)
	}

	dbTask, err := f.persist.FetchWorkerTask(ctx, dbWorker)
	if err != nil {
		logger.Error().Err(err).Msg("fetching task assigned to worker")
		return sendAPIError(e, http.StatusInternalServerError, "error fetching task assigned to worker: %v", err)
	}

	logger.Debug().Msg("fetched worker")
	apiWorker := workerDBtoAPI(*dbWorker)

	if dbTask != nil {
		apiWorkerTask := api.WorkerTask{
			TaskSummary: taskDBtoSummary(dbTask),
			JobId:       dbTask.Job.UUID,
		}
		apiWorker.Task = &apiWorkerTask
	}

	return e.JSON(http.StatusOK, apiWorker)
}

func (f *Flamenco) DeleteWorker(e echo.Context, workerUUID string) error {
	logger := requestLogger(e)
	logger = logger.With().Str("worker", workerUUID).Logger()

	if !uuid.IsValid(workerUUID) {
		return sendAPIError(e, http.StatusBadRequest, "not a valid UUID")
	}

	// All information to do the deletion is known, so even when the client
	// disconnects, the deletion should be completed.
	ctx, ctxCancel := bgContext()
	defer ctxCancel()

	// Fetch the worker in order to re-queue its tasks.
	worker, err := f.persist.FetchWorker(ctx, workerUUID)
	if errors.Is(err, persistence.ErrWorkerNotFound) {
		logger.Debug().Msg("deletion of non-existent worker requested")
		return sendAPIError(e, http.StatusNotFound, "worker %q not found", workerUUID)
	}
	if err != nil {
		logger.Error().Err(err).Msg("fetching worker for deletion")
		return sendAPIError(e, http.StatusInternalServerError,
			"error fetching worker for deletion: %v", err)
	}

	err = f.stateMachine.RequeueActiveTasksOfWorker(ctx, worker, "worker is being deleted")
	if err != nil {
		logger.Error().Err(err).Msg("requeueing tasks before deleting worker")
		return sendAPIError(e, http.StatusInternalServerError,
			"error requeueing tasks before deleting worker: %v", err)
	}

	// Actually delete the worker.
	err = f.persist.DeleteWorker(ctx, workerUUID)
	if errors.Is(err, persistence.ErrWorkerNotFound) {
		logger.Debug().Msg("deletion of non-existent worker requested")
		return sendAPIError(e, http.StatusNotFound, "worker %q not found", workerUUID)
	}
	if err != nil {
		logger.Error().Err(err).Msg("deleting worker")
		return sendAPIError(e, http.StatusInternalServerError, "error deleting worker: %v", err)
	}
	logger.Info().Msg("deleted worker")

	// It would be cleaner to re-fetch the Worker from the database and get the
	// exact 'deleted at' timestamp from there, but that would require more DB
	// operations, and this is accurate enough for a quick broadcast via SocketIO.
	now := f.clock.Now()

	// Broadcast the fact that this worker was just deleted.
	update := webupdates.NewWorkerUpdate(worker)
	update.DeletedAt = &now
	f.broadcaster.BroadcastWorkerUpdate(update)

	return e.NoContent(http.StatusNoContent)
}

func (f *Flamenco) RequestWorkerStatusChange(e echo.Context, workerUUID string) error {
	logger := requestLogger(e)
	logger = logger.With().Str("worker", workerUUID).Logger()

	if !uuid.IsValid(workerUUID) {
		return sendAPIError(e, http.StatusBadRequest, "not a valid UUID")
	}

	// Decode the request body.
	var change api.WorkerStatusChangeRequest
	if err := e.Bind(&change); err != nil {
		logger.Warn().Err(err).Msg("bad request received")
		return sendAPIError(e, http.StatusBadRequest, "invalid format")
	}

	// Fetch the worker.
	dbWorker, err := f.persist.FetchWorker(e.Request().Context(), workerUUID)
	if errors.Is(err, persistence.ErrWorkerNotFound) {
		logger.Debug().Msg("non-existent worker requested")
		return sendAPIError(e, http.StatusNotFound, "worker %q not found", workerUUID)
	}
	if err != nil {
		logger.Error().Err(err).Msg("fetching worker")
		return sendAPIError(e, http.StatusInternalServerError, "error fetching worker: %v", err)
	}

	logger = logger.With().
		Str("status", string(dbWorker.Status)).
		Str("requested", string(change.Status)).
		Bool("lazy", change.IsLazy).
		Logger()

	if change.Status == api.WorkerStatusRestart && !dbWorker.CanRestart {
		logger.Error().Msg("worker cannot be restarted, rejecting status change request")
		return sendAPIError(e, http.StatusPreconditionFailed,
			"worker %q does not know how to restart", workerUUID)
	}

	logger.Info().Msg("worker status change requested")

	if dbWorker.Status == change.Status {
		// Requesting that the worker should go to its current status basically
		// means cancelling any previous status change request.
		dbWorker.StatusChangeClear()
	} else {
		dbWorker.StatusChangeRequest(change.Status, change.IsLazy)
	}

	// Store the status change.
	if err := f.persist.SaveWorker(e.Request().Context(), dbWorker); err != nil {
		logger.Error().Err(err).Msg("saving worker after status change request")
		return sendAPIError(e, http.StatusInternalServerError, "error saving worker: %v", err)
	}

	// Broadcast the change.
	update := webupdates.NewWorkerUpdate(dbWorker)
	f.broadcaster.BroadcastWorkerUpdate(update)

	return e.NoContent(http.StatusNoContent)
}

func (f *Flamenco) SetWorkerTags(e echo.Context, workerUUID string) error {
	ctx := e.Request().Context()
	logger := requestLogger(e)
	logger = logger.With().Str("worker", workerUUID).Logger()

	if !uuid.IsValid(workerUUID) {
		return sendAPIError(e, http.StatusBadRequest, "not a valid UUID")
	}

	// Decode the request body.
	var change api.WorkerTagChangeRequest
	if err := e.Bind(&change); err != nil {
		logger.Warn().Err(err).Msg("bad request received")
		return sendAPIError(e, http.StatusBadRequest, "invalid format")
	}

	// Fetch the worker.
	dbWorker, err := f.persist.FetchWorker(ctx, workerUUID)
	if errors.Is(err, persistence.ErrWorkerNotFound) {
		logger.Debug().Msg("non-existent worker requested")
		return sendAPIError(e, http.StatusNotFound, "worker %q not found", workerUUID)
	}
	if err != nil {
		logger.Error().Err(err).Msg("fetching worker")
		return sendAPIError(e, http.StatusInternalServerError, "error fetching worker: %v", err)
	}

	logger = logger.With().
		Strs("tags", change.TagIds).
		Logger()
	logger.Info().Msg("worker tag change requested")

	// Store the new tag assignment.
	if err := f.persist.WorkerSetTags(ctx, dbWorker, change.TagIds); err != nil {
		logger.Error().Err(err).Msg("saving worker after tag change request")
		return sendAPIError(e, http.StatusInternalServerError, "error saving worker: %v", err)
	}

	// Broadcast the change.
	update := webupdates.NewWorkerUpdate(dbWorker)
	f.broadcaster.BroadcastWorkerUpdate(update)

	return e.NoContent(http.StatusNoContent)
}

func (f *Flamenco) DeleteWorkerTag(e echo.Context, tagUUID string) error {
	ctx := e.Request().Context()
	logger := requestLogger(e)
	logger = logger.With().Str("tag", tagUUID).Logger()

	if !uuid.IsValid(tagUUID) {
		return sendAPIError(e, http.StatusBadRequest, "not a valid UUID")
	}

	err := f.persist.DeleteWorkerTag(ctx, tagUUID)
	switch {
	case errors.Is(err, persistence.ErrWorkerTagNotFound):
		logger.Debug().Msg("non-existent worker tag requested")
		return sendAPIError(e, http.StatusNotFound, "worker tag %q not found", tagUUID)
	case err != nil:
		logger.Error().Err(err).Msg("deleting worker tag")
		return sendAPIError(e, http.StatusInternalServerError, "error deleting worker tag: %v", err)
	}

	// TODO: SocketIO broadcast of tag deletion.

	logger.Info().Msg("worker tag deleted")
	return e.NoContent(http.StatusNoContent)
}

func (f *Flamenco) FetchWorkerTag(e echo.Context, tagUUID string) error {
	ctx := e.Request().Context()
	logger := requestLogger(e)
	logger = logger.With().Str("tag", tagUUID).Logger()

	if !uuid.IsValid(tagUUID) {
		return sendAPIError(e, http.StatusBadRequest, "not a valid UUID")
	}

	tag, err := f.persist.FetchWorkerTag(ctx, tagUUID)
	switch {
	case errors.Is(err, persistence.ErrWorkerTagNotFound):
		logger.Debug().Msg("non-existent worker tag requested")
		return sendAPIError(e, http.StatusNotFound, "worker tag %q not found", tagUUID)
	case err != nil:
		logger.Error().Err(err).Msg("fetching worker tag")
		return sendAPIError(e, http.StatusInternalServerError, "error fetching worker tag: %v", err)
	}

	return e.JSON(http.StatusOK, workerTagDBtoAPI(*tag))
}

func (f *Flamenco) UpdateWorkerTag(e echo.Context, tagUUID string) error {
	ctx := e.Request().Context()
	logger := requestLogger(e)
	logger = logger.With().Str("tag", tagUUID).Logger()

	if !uuid.IsValid(tagUUID) {
		return sendAPIError(e, http.StatusBadRequest, "not a valid UUID")
	}

	// Decode the request body.
	var update api.UpdateWorkerTagJSONBody
	if err := e.Bind(&update); err != nil {
		logger.Warn().Err(err).Msg("bad request received")
		return sendAPIError(e, http.StatusBadRequest, "invalid format")
	}

	dbTag, err := f.persist.FetchWorkerTag(ctx, tagUUID)
	switch {
	case errors.Is(err, persistence.ErrWorkerTagNotFound):
		logger.Debug().Msg("non-existent worker tag requested")
		return sendAPIError(e, http.StatusNotFound, "worker tag %q not found", tagUUID)
	case err != nil:
		logger.Error().Err(err).Msg("fetching worker tag")
		return sendAPIError(e, http.StatusInternalServerError, "error fetching worker tag: %v", err)
	}

	// Update the tag.
	dbTag.Name = update.Name
	if update.Description == nil {
		dbTag.Description = ""
	} else {
		dbTag.Description = *update.Description
	}

	if err := f.persist.SaveWorkerTag(ctx, dbTag); err != nil {
		logger.Error().Err(err).Msg("saving worker tag")
		return sendAPIError(e, http.StatusInternalServerError, "error saving worker tag")
	}

	// TODO: SocketIO broadcast of tag update.

	return e.NoContent(http.StatusNoContent)
}

func (f *Flamenco) FetchWorkerTags(e echo.Context) error {
	ctx := e.Request().Context()
	logger := requestLogger(e)

	dbTags, err := f.persist.FetchWorkerTags(ctx)
	if err != nil {
		logger.Error().Err(err).Msg("fetching worker tags")
		return sendAPIError(e, http.StatusInternalServerError, "error saving worker tag")
	}

	apiTags := []api.WorkerTag{}
	for _, dbTag := range dbTags {
		apiTag := workerTagDBtoAPI(*dbTag)
		apiTags = append(apiTags, apiTag)
	}

	tagList := api.WorkerTagList{
		Tags: &apiTags,
	}
	return e.JSON(http.StatusOK, &tagList)
}

func (f *Flamenco) CreateWorkerTag(e echo.Context) error {
	ctx := e.Request().Context()
	logger := requestLogger(e)

	// Decode the request body.
	var apiTag api.CreateWorkerTagJSONBody
	if err := e.Bind(&apiTag); err != nil {
		logger.Warn().Err(err).Msg("bad request received")
		return sendAPIError(e, http.StatusBadRequest, "invalid format")
	}

	// Convert to persistence layer model.
	var tagUUID string
	if apiTag.Id != nil && *apiTag.Id != "" {
		tagUUID = *apiTag.Id
	} else {
		tagUUID = uuid.New()
	}

	dbTag := persistence.WorkerTag{
		UUID: tagUUID,
		Name: apiTag.Name,
	}
	if apiTag.Description != nil {
		dbTag.Description = *apiTag.Description
	}

	// Store in the database.
	if err := f.persist.CreateWorkerTag(ctx, &dbTag); err != nil {
		logger.Error().Err(err).Msg("creating worker tag")
		return sendAPIError(e, http.StatusInternalServerError, "error creating worker tag")
	}

	// TODO: SocketIO broadcast of tag creation.

	return e.JSON(http.StatusOK, workerTagDBtoAPI(dbTag))
}

func workerSummary(w persistence.Worker) api.WorkerSummary {
	summary := api.WorkerSummary{
		Id:         w.UUID,
		Name:       w.Name,
		Status:     w.Status,
		Version:    w.Software,
		CanRestart: w.CanRestart,
	}
	if w.StatusRequested != "" {
		summary.StatusChange = &api.WorkerStatusChangeRequest{
			Status: w.StatusRequested,
			IsLazy: w.LazyStatusRequest,
		}
	}

	if !w.LastSeenAt.IsZero() {
		summary.LastSeen = &w.LastSeenAt
	}

	return summary
}

func workerDBtoAPI(w persistence.Worker) api.Worker {
	apiWorker := api.Worker{
		WorkerSummary:      workerSummary(w),
		IpAddress:          w.Address,
		Platform:           w.Platform,
		SupportedTaskTypes: w.TaskTypes(),
	}

	if len(w.Tags) > 0 {
		tags := []api.WorkerTag{}
		for i := range w.Tags {
			tags = append(tags, workerTagDBtoAPI(*w.Tags[i]))
		}
		apiWorker.Tags = &tags
	}

	return apiWorker
}

func workerTagDBtoAPI(wc persistence.WorkerTag) api.WorkerTag {
	uuid := wc.UUID // Take a copy for safety.

	apiTag := api.WorkerTag{
		Id:   &uuid,
		Name: wc.Name,
	}
	if len(wc.Description) > 0 {
		apiTag.Description = &wc.Description
	}
	return apiTag
}
