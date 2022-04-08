// SPDX-License-Identifier: GPL-3.0-or-later
package webupdates

import (
	"github.com/rs/zerolog/log"

	"git.blender.org/flamenco/internal/manager/persistence"
	"git.blender.org/flamenco/pkg/api"
)

// NewJobUpdate returns a partial JobUpdate struct for the given job.
// It only fills in the fields that represent the current state of the job. For
// example, it omits `PreviousStatus`. The ommitted fields can be filled in by
// the caller.
func NewJobUpdate(job *persistence.Job) api.JobUpdate {
	jobUpdate := api.JobUpdate{
		Id:      job.UUID,
		Name:    &job.Name,
		Updated: job.UpdatedAt,
		Status:  job.Status,
	}
	return jobUpdate
}

// BroadcastJobUpdate sends the job update to clients.
func (b *BiDirComms) BroadcastJobUpdate(jobUpdate api.JobUpdate) {
	log.Debug().Interface("jobUpdate", jobUpdate).Msg("socketIO: broadcasting job update")
	b.BroadcastTo(SocketIORoomJobs, SIOEventJobUpdate, jobUpdate)
}

// BroadcastNewJob sends a "new job" notification to clients.
func (b *BiDirComms) BroadcastNewJob(jobUpdate api.JobUpdate) {
	if jobUpdate.PreviousStatus != nil {
		log.Warn().Interface("jobUpdate", jobUpdate).Msg("socketIO: new jobs should not have a previous state")
		jobUpdate.PreviousStatus = nil
	}

	log.Debug().Interface("jobUpdate", jobUpdate).Msg("socketIO: broadcasting new job")
	b.BroadcastTo(SocketIORoomJobs, SIOEventJobUpdate, jobUpdate)
}