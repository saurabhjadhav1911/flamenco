package webupdates

// SPDX-License-Identifier: GPL-3.0-or-later

import (
	"github.com/rs/zerolog/log"
	"projects.blender.org/studio/flamenco/internal/manager/persistence"
	"projects.blender.org/studio/flamenco/pkg/api"
)

// NewWorkerTagUpdate returns a partial SocketIOWorkerTagUpdate struct for the
// given worker tag. It only fills in the fields that represent the current
// state of the tag.
func NewWorkerTagUpdate(tag *persistence.WorkerTag) api.SocketIOWorkerTagUpdate {
	tagUpdate := api.SocketIOWorkerTagUpdate{
		Tag: api.WorkerTag{
			Id:          &tag.UUID,
			Name:        tag.Name,
			Description: &tag.Description,
		},
	}
	return tagUpdate
}

// NewWorkerTagDeletedUpdate returns a SocketIOWorkerTagUpdate struct that indicates
// the worker tag has been deleted.
func NewWorkerTagDeletedUpdate(tagUUID string) api.SocketIOWorkerTagUpdate {
	wasDeleted := true
	tagUpdate := api.SocketIOWorkerTagUpdate{
		Tag: api.WorkerTag{
			Id: &tagUUID,
		},
		WasDeleted: &wasDeleted,
	}
	return tagUpdate
}

// BroadcastWorkerTagUpdate sends the worker tag update to clients.
func (b *BiDirComms) BroadcastWorkerTagUpdate(WorkerTagUpdate api.SocketIOWorkerTagUpdate) {
	log.Debug().Interface("WorkerTagUpdate", WorkerTagUpdate).Msg("socketIO: broadcasting worker tag update")
	b.BroadcastTo(SocketIORoomWorkerTags, SIOEventWorkerTagUpdate, WorkerTagUpdate)
}

// BroadcastNewWorkerTag sends a "new worker tag" notification to clients.
func (b *BiDirComms) BroadcastNewWorkerTag(WorkerTagUpdate api.SocketIOWorkerTagUpdate) {
	log.Debug().Interface("WorkerTagUpdate", WorkerTagUpdate).Msg("socketIO: broadcasting new worker tag")
	b.BroadcastTo(SocketIORoomWorkerTags, SIOEventWorkerTagUpdate, WorkerTagUpdate)
}
