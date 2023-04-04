package api_impl

// SPDX-License-Identifier: GPL-3.0-or-later

import (
	"errors"
	"net/http"

	"git.blender.org/flamenco/internal/manager/persistence"
	"git.blender.org/flamenco/internal/manager/webupdates"
	"git.blender.org/flamenco/internal/uuid"
	"git.blender.org/flamenco/pkg/api"
	"github.com/labstack/echo/v4"
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

func (f *Flamenco) SetWorkerClusters(e echo.Context, workerUUID string) error {
	ctx := e.Request().Context()
	logger := requestLogger(e)
	logger = logger.With().Str("worker", workerUUID).Logger()

	if !uuid.IsValid(workerUUID) {
		return sendAPIError(e, http.StatusBadRequest, "not a valid UUID")
	}

	// Decode the request body.
	var change api.WorkerClusterChangeRequest
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
		Strs("clusters", change.ClusterIds).
		Logger()
	logger.Info().Msg("worker cluster change requested")

	// Store the new cluster assignment.
	if err := f.persist.WorkerSetClusters(ctx, dbWorker, change.ClusterIds); err != nil {
		logger.Error().Err(err).Msg("saving worker after cluster change request")
		return sendAPIError(e, http.StatusInternalServerError, "error saving worker: %v", err)
	}

	// Broadcast the change.
	update := webupdates.NewWorkerUpdate(dbWorker)
	f.broadcaster.BroadcastWorkerUpdate(update)

	return e.NoContent(http.StatusNoContent)
}

func (f *Flamenco) DeleteWorkerCluster(e echo.Context, clusterUUID string) error {
	ctx := e.Request().Context()
	logger := requestLogger(e)
	logger = logger.With().Str("cluster", clusterUUID).Logger()

	if !uuid.IsValid(clusterUUID) {
		return sendAPIError(e, http.StatusBadRequest, "not a valid UUID")
	}

	err := f.persist.DeleteWorkerCluster(ctx, clusterUUID)
	switch {
	case errors.Is(err, persistence.ErrWorkerClusterNotFound):
		logger.Debug().Msg("non-existent worker cluster requested")
		return sendAPIError(e, http.StatusNotFound, "worker cluster %q not found", clusterUUID)
	case err != nil:
		logger.Error().Err(err).Msg("deleting worker cluster")
		return sendAPIError(e, http.StatusInternalServerError, "error deleting worker cluster: %v", err)
	}

	// TODO: SocketIO broadcast of cluster deletion.

	logger.Info().Msg("worker cluster deleted")
	return e.NoContent(http.StatusNoContent)
}

func (f *Flamenco) FetchWorkerCluster(e echo.Context, clusterUUID string) error {
	ctx := e.Request().Context()
	logger := requestLogger(e)
	logger = logger.With().Str("cluster", clusterUUID).Logger()

	if !uuid.IsValid(clusterUUID) {
		return sendAPIError(e, http.StatusBadRequest, "not a valid UUID")
	}

	cluster, err := f.persist.FetchWorkerCluster(ctx, clusterUUID)
	switch {
	case errors.Is(err, persistence.ErrWorkerClusterNotFound):
		logger.Debug().Msg("non-existent worker cluster requested")
		return sendAPIError(e, http.StatusNotFound, "worker cluster %q not found", clusterUUID)
	case err != nil:
		logger.Error().Err(err).Msg("fetching worker cluster")
		return sendAPIError(e, http.StatusInternalServerError, "error fetching worker cluster: %v", err)
	}

	return e.JSON(http.StatusOK, workerClusterDBtoAPI(*cluster))
}

func (f *Flamenco) UpdateWorkerCluster(e echo.Context, clusterUUID string) error {
	ctx := e.Request().Context()
	logger := requestLogger(e)
	logger = logger.With().Str("cluster", clusterUUID).Logger()

	if !uuid.IsValid(clusterUUID) {
		return sendAPIError(e, http.StatusBadRequest, "not a valid UUID")
	}

	// Decode the request body.
	var update api.UpdateWorkerClusterJSONBody
	if err := e.Bind(&update); err != nil {
		logger.Warn().Err(err).Msg("bad request received")
		return sendAPIError(e, http.StatusBadRequest, "invalid format")
	}

	dbCluster, err := f.persist.FetchWorkerCluster(ctx, clusterUUID)
	switch {
	case errors.Is(err, persistence.ErrWorkerClusterNotFound):
		logger.Debug().Msg("non-existent worker cluster requested")
		return sendAPIError(e, http.StatusNotFound, "worker cluster %q not found", clusterUUID)
	case err != nil:
		logger.Error().Err(err).Msg("fetching worker cluster")
		return sendAPIError(e, http.StatusInternalServerError, "error fetching worker cluster: %v", err)
	}

	// Update the cluster.
	dbCluster.Name = update.Name
	if update.Description == nil {
		dbCluster.Description = ""
	} else {
		dbCluster.Description = *update.Description
	}

	if err := f.persist.SaveWorkerCluster(ctx, dbCluster); err != nil {
		logger.Error().Err(err).Msg("saving worker cluster")
		return sendAPIError(e, http.StatusInternalServerError, "error saving worker cluster")
	}

	// TODO: SocketIO broadcast of cluster update.

	return e.NoContent(http.StatusNoContent)
}

func (f *Flamenco) FetchWorkerClusters(e echo.Context) error {
	ctx := e.Request().Context()
	logger := requestLogger(e)

	dbClusters, err := f.persist.FetchWorkerClusters(ctx)
	if err != nil {
		logger.Error().Err(err).Msg("fetching worker clusters")
		return sendAPIError(e, http.StatusInternalServerError, "error saving worker cluster")
	}

	apiClusters := []api.WorkerCluster{}
	for _, dbCluster := range dbClusters {
		apiCluster := workerClusterDBtoAPI(*dbCluster)
		apiClusters = append(apiClusters, apiCluster)
	}

	clusterList := api.WorkerClusterList{
		Clusters: &apiClusters,
	}
	return e.JSON(http.StatusOK, &clusterList)
}

func (f *Flamenco) CreateWorkerCluster(e echo.Context) error {
	ctx := e.Request().Context()
	logger := requestLogger(e)

	// Decode the request body.
	var apiCluster api.CreateWorkerClusterJSONBody
	if err := e.Bind(&apiCluster); err != nil {
		logger.Warn().Err(err).Msg("bad request received")
		return sendAPIError(e, http.StatusBadRequest, "invalid format")
	}

	// Convert to persistence layer model.
	var clusterUUID string
	if apiCluster.Id != nil && *apiCluster.Id != "" {
		clusterUUID = *apiCluster.Id
	} else {
		clusterUUID = uuid.New()
	}

	dbCluster := persistence.WorkerCluster{
		UUID: clusterUUID,
		Name: apiCluster.Name,
	}
	if apiCluster.Description != nil {
		dbCluster.Description = *apiCluster.Description
	}

	// Store in the database.
	if err := f.persist.CreateWorkerCluster(ctx, &dbCluster); err != nil {
		logger.Error().Err(err).Msg("creating worker cluster")
		return sendAPIError(e, http.StatusInternalServerError, "error creating worker cluster")
	}

	// TODO: SocketIO broadcast of cluster creation.

	return e.JSON(http.StatusOK, workerClusterDBtoAPI(dbCluster))
}

func workerSummary(w persistence.Worker) api.WorkerSummary {
	summary := api.WorkerSummary{
		Id:      w.UUID,
		Name:    w.Name,
		Status:  w.Status,
		Version: w.Software,
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

	if len(w.Clusters) > 0 {
		clusters := []api.WorkerCluster{}
		for i := range w.Clusters {
			clusters = append(clusters, workerClusterDBtoAPI(*w.Clusters[i]))
		}
		apiWorker.Clusters = &clusters
	}

	return apiWorker
}

func workerClusterDBtoAPI(wc persistence.WorkerCluster) api.WorkerCluster {
	uuid := wc.UUID // Take a copy for safety.

	apiCluster := api.WorkerCluster{
		Id:   &uuid,
		Name: wc.Name,
	}
	if len(wc.Description) > 0 {
		apiCluster.Description = &wc.Description
	}
	return apiCluster
}
