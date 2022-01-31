package api_impl

/* ***** BEGIN GPL LICENSE BLOCK *****
 *
 * Original Code Copyright (C) 2022 Blender Foundation.
 *
 * This file is part of Flamenco.
 *
 * Flamenco is free software: you can redistribute it and/or modify it under
 * the terms of the GNU General Public License as published by the Free Software
 * Foundation, either version 3 of the License, or (at your option) any later
 * version.
 *
 * Flamenco is distributed in the hope that it will be useful, but WITHOUT ANY
 * WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR
 * A PARTICULAR PURPOSE.  See the GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License along with
 * Flamenco.  If not, see <https://www.gnu.org/licenses/>.
 *
 * ***** END GPL LICENSE BLOCK ***** */

import (
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"

	"gitlab.com/blender/flamenco-ng-poc/internal/manager/persistence"
	"gitlab.com/blender/flamenco-ng-poc/pkg/api"
)

// RegisterWorker registers a new worker and stores it in the database.
func (f *Flamenco) RegisterWorker(e echo.Context) error {
	logger := requestLogger(e)

	var req api.RegisterWorkerJSONBody
	err := e.Bind(&req)
	if err != nil {
		logger.Warn().Err(err).Msg("bad request received")
		return sendAPIError(e, http.StatusBadRequest, "invalid format")
	}

	// TODO: validate the request, should at least have non-empty name, secret, and platform.

	logger.Info().Str("nickname", req.Nickname).Msg("registering new worker")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Secret), bcrypt.DefaultCost)
	if err != nil {
		logger.Warn().Err(err).Msg("error hashing worker password")
		return sendAPIError(e, http.StatusBadRequest, "error hashing password")
	}

	dbWorker := persistence.Worker{
		UUID:               uuid.New().String(),
		Name:               req.Nickname,
		Secret:             string(hashedPassword),
		Platform:           req.Platform,
		Address:            e.RealIP(),
		SupportedTaskTypes: strings.Join(req.SupportedTaskTypes, ","),
	}
	if err := f.persist.CreateWorker(e.Request().Context(), &dbWorker); err != nil {
		logger.Warn().Err(err).Msg("error creating new worker in DB")
		return sendAPIError(e, http.StatusBadRequest, "error registering worker")
	}

	return e.JSON(http.StatusOK, &api.RegisteredWorker{
		Uuid:               dbWorker.UUID,
		Nickname:           dbWorker.Name,
		Address:            dbWorker.Address,
		LastActivity:       dbWorker.LastActivity,
		Platform:           dbWorker.Platform,
		Software:           dbWorker.Software,
		Status:             dbWorker.Status,
		SupportedTaskTypes: strings.Split(dbWorker.SupportedTaskTypes, ","),
	})
}

func (f *Flamenco) SignOn(e echo.Context) error {
	logger := requestLogger(e)

	var req api.SignOnJSONBody
	err := e.Bind(&req)
	if err != nil {
		logger.Warn().Err(err).Msg("bad request received")
		return sendAPIError(e, http.StatusBadRequest, "invalid format")
	}

	logger.Info().Msg("worker signing on")

	w := requestWorkerOrPanic(e)
	w.Status = api.WorkerStatusStarting
	err = f.persist.SaveWorker(e.Request().Context(), w)
	if err != nil {
		logger.Warn().Err(err).
			Str("newStatus", string(w.Status)).
			Msg("error storing Worker in database")
		return sendAPIError(e, http.StatusInternalServerError, "error storing worker in database")
	}

	resp := api.WorkerStateChange{}
	if w.StatusRequested != "" {
		resp.StatusRequested = w.StatusRequested
	} else {
		resp.StatusRequested = api.WorkerStatusAwake
	}
	return e.JSON(http.StatusOK, resp)
}

func (f *Flamenco) SignOff(e echo.Context) error {
	logger := requestLogger(e)

	var req api.SignOnJSONBody
	err := e.Bind(&req)
	if err != nil {
		logger.Warn().Err(err).Msg("bad request received")
		return sendAPIError(e, http.StatusBadRequest, "invalid format")
	}

	logger.Info().Msg("worker signing off")
	w := requestWorkerOrPanic(e)
	w.Status = api.WorkerStatusOffline
	// TODO: check whether we should pass the request context here, or a generic
	// background context, as this should be stored even when the HTTP connection
	// is aborted.
	err = f.persist.SaveWorker(e.Request().Context(), w)
	if err != nil {
		logger.Warn().
			Err(err).
			Str("newStatus", string(w.Status)).
			Msg("error storing worker status in database")
		return sendAPIError(e, http.StatusInternalServerError, "error storing new status in database")
	}

	return e.String(http.StatusNoContent, "")
}

// (GET /api/worker/state)
func (f *Flamenco) WorkerState(e echo.Context) error {
	// TODO: look up proper status in DB.
	return e.String(http.StatusNoContent, "")
}

// Worker changed state. This could be as acknowledgement of a Manager-requested state change, or in response to worker-local signals.
// (POST /api/worker/state-changed)
func (f *Flamenco) WorkerStateChanged(e echo.Context) error {
	logger := requestLogger(e)

	var req api.WorkerStateChangedJSONRequestBody
	err := e.Bind(&req)
	if err != nil {
		logger.Warn().Err(err).Msg("bad request received")
		return sendAPIError(e, http.StatusBadRequest, "invalid format")
	}

	logger.Info().Str("newStatus", string(req.Status)).Msg("worker changed status")

	w := requestWorkerOrPanic(e)
	w.Status = req.Status
	err = f.persist.SaveWorker(e.Request().Context(), w)
	if err != nil {
		logger.Warn().Err(err).
			Str("newStatus", string(w.Status)).
			Msg("error storing Worker in database")
		return sendAPIError(e, http.StatusInternalServerError, "error storing worker in database")
	}

	return e.String(http.StatusNoContent, "")
}

func (f *Flamenco) ScheduleTask(e echo.Context) error {
	logger := requestLogger(e)
	logger.Info().Msg("worker requesting task")

	return e.JSON(http.StatusOK, &api.AssignedTask{
		Uuid: uuid.New().String(),
		Commands: []api.Command{
			{Name: "echo", Settings: echo.Map{"payload": "Simon says \"Shaders!\""}},
			{Name: "blender", Settings: echo.Map{"blender_cmd": "/shared/bin/blender"}},
		},
		Job:         uuid.New().String(),
		JobPriority: 50,
		JobType:     "blender-render",
		Name:        "A1032",
		Priority:    50,
		Status:      "active",
		TaskType:    "blender-render",
		User:        "",
	})
}
