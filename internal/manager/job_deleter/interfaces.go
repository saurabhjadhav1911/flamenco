package job_deleter

// SPDX-License-Identifier: GPL-3.0-or-later

import (
	"context"
	"time"

	"projects.blender.org/studio/flamenco/internal/manager/local_storage"
	"projects.blender.org/studio/flamenco/internal/manager/persistence"
	"projects.blender.org/studio/flamenco/internal/manager/webupdates"
	"projects.blender.org/studio/flamenco/pkg/api"
	"projects.blender.org/studio/flamenco/pkg/shaman"
)

// Generate mock implementations of these interfaces.
//go:generate go run github.com/golang/mock/mockgen -destination mocks/interfaces_mock.gen.go -package mocks projects.blender.org/studio/flamenco/internal/manager/job_deleter PersistenceService,Storage,ChangeBroadcaster,Shaman

type PersistenceService interface {
	FetchJob(ctx context.Context, jobUUID string) (*persistence.Job, error)

	RequestJobDeletion(ctx context.Context, j *persistence.Job) error
	RequestJobMassDeletion(ctx context.Context, lastUpdatedMax time.Time) ([]string, error)

	// FetchJobsDeletionRequested returns the UUIDs of to-be-deleted jobs.
	FetchJobsDeletionRequested(ctx context.Context) ([]string, error)
	DeleteJob(ctx context.Context, jobUUID string) error
}

// PersistenceService should be a subset of persistence.DB
var _ PersistenceService = (*persistence.DB)(nil)

type Storage interface {
	// RemoveJobStorage removes from disk the directory for storing job-related files.
	RemoveJobStorage(ctx context.Context, jobUUID string) error
}

var _ Storage = (*local_storage.StorageInfo)(nil)

type ChangeBroadcaster interface {
	// BroadcastJobUpdate sends the job update to SocketIO clients.
	BroadcastJobUpdate(jobUpdate api.SocketIOJobUpdate)
}

// ChangeBroadcaster should be a subset of webupdates.BiDirComms
var _ ChangeBroadcaster = (*webupdates.BiDirComms)(nil)

type Shaman interface {
	// IsEnabled returns whether this Shaman service is enabled or not.
	IsEnabled() bool

	// EraseCheckout deletes the symlinks and the directory structure that makes up the checkout.
	EraseCheckout(checkoutID string) error
}

var _ Shaman = (*shaman.Server)(nil)
