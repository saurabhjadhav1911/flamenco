package persistence

// SPDX-License-Identifier: GPL-3.0-or-later

import (
	"testing"
	"time"

	"git.blender.org/flamenco/internal/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateFetchCluster(t *testing.T) {
	f := workerTestFixtures(t, 1*time.Second)
	defer f.done()

	// Test fetching non-existent cluster
	fetchedCluster, err := f.db.FetchWorkerCluster(f.ctx, "7ee21bc8-ff1a-42d2-a6b6-cc4b529b189f")
	assert.ErrorIs(t, err, ErrWorkerClusterNotFound)
	assert.Nil(t, fetchedCluster)

	// New cluster creation is already done in the workerTestFixtures() call.
	assert.NotNil(t, f.cluster)

	fetchedCluster, err = f.db.FetchWorkerCluster(f.ctx, f.cluster.UUID)
	require.NoError(t, err)
	assert.NotNil(t, fetchedCluster)

	// Test contents of fetched cluster.
	assert.Equal(t, f.cluster.UUID, fetchedCluster.UUID)
	assert.Equal(t, f.cluster.Name, fetchedCluster.Name)
	assert.Equal(t, f.cluster.Description, fetchedCluster.Description)
	assert.Zero(t, fetchedCluster.Workers)
}

func TestFetchDeleteClusters(t *testing.T) {
	f := workerTestFixtures(t, 1*time.Second)
	defer f.done()

	secondCluster := WorkerCluster{
		UUID:        uuid.New(),
		Name:        "arbeiderscluster",
		Description: "Worker cluster in Dutch",
	}

	require.NoError(t, f.db.CreateWorkerCluster(f.ctx, &secondCluster))

	allClusters, err := f.db.FetchWorkerClusters(f.ctx)
	require.NoError(t, err)

	require.Len(t, allClusters, 2)
	var allClusterIDs [2]string
	for idx := range allClusters {
		allClusterIDs[idx] = allClusters[idx].UUID
	}
	assert.Contains(t, allClusterIDs, f.cluster.UUID)
	assert.Contains(t, allClusterIDs, secondCluster.UUID)

	// Test deleting the 2nd cluster.
	require.NoError(t, f.db.DeleteWorkerCluster(f.ctx, secondCluster.UUID))

	allClusters, err = f.db.FetchWorkerClusters(f.ctx)
	require.NoError(t, err)
	require.Len(t, allClusters, 1)
	assert.Equal(t, f.cluster.UUID, allClusters[0].UUID)
}

func TestAssignUnassignWorkerClusters(t *testing.T) {
	f := workerTestFixtures(t, 1*time.Second)
	defer f.done()

	assertClusters := func(msgLabel string, clusterUUIDs ...string) {
		w, err := f.db.FetchWorker(f.ctx, f.worker.UUID)
		require.NoError(t, err)

		// Catch doubly-reported clusters, as the maps below would hide those cases.
		assert.Len(t, w.Clusters, len(clusterUUIDs), msgLabel)

		expectClusters := make(map[string]bool)
		for _, cid := range clusterUUIDs {
			expectClusters[cid] = true
		}

		actualClusters := make(map[string]bool)
		for _, c := range w.Clusters {
			actualClusters[c.UUID] = true
		}

		assert.Equal(t, expectClusters, actualClusters, msgLabel)
	}

	secondCluster := WorkerCluster{
		UUID:        uuid.New(),
		Name:        "arbeiderscluster",
		Description: "Worker cluster in Dutch",
	}

	require.NoError(t, f.db.CreateWorkerCluster(f.ctx, &secondCluster))

	// By default the Worker should not be part of a cluster.
	assertClusters("default cluster assignment")

	require.NoError(t, f.db.WorkerSetClusters(f.ctx, f.worker, []string{f.cluster.UUID}))
	assertClusters("setting one cluster", f.cluster.UUID)

	// Double assignments should also just work.
	require.NoError(t, f.db.WorkerSetClusters(f.ctx, f.worker, []string{f.cluster.UUID, f.cluster.UUID}))
	assertClusters("setting twice the same cluster", f.cluster.UUID)

	// Multiple cluster memberships.
	require.NoError(t, f.db.WorkerSetClusters(f.ctx, f.worker, []string{f.cluster.UUID, secondCluster.UUID}))
	assertClusters("setting two different clusters", f.cluster.UUID, secondCluster.UUID)

	// Remove memberships.
	require.NoError(t, f.db.WorkerSetClusters(f.ctx, f.worker, []string{secondCluster.UUID}))
	assertClusters("unassigning from first cluster", secondCluster.UUID)
	require.NoError(t, f.db.WorkerSetClusters(f.ctx, f.worker, []string{}))
	assertClusters("unassigning from second cluster")
}

func TestSaveWorkerCluster(t *testing.T) {
	f := workerTestFixtures(t, 1*time.Second)
	defer f.done()

	f.cluster.Name = "übercluster"
	f.cluster.Description = "ʻO kēlā hui ma laila"
	require.NoError(t, f.db.SaveWorkerCluster(f.ctx, f.cluster))

	fetched, err := f.db.FetchWorkerCluster(f.ctx, f.cluster.UUID)
	require.NoError(t, err)
	assert.Equal(t, f.cluster.Name, fetched.Name)
	assert.Equal(t, f.cluster.Description, fetched.Description)
}

func TestDeleteWorkerClusterWithWorkersAssigned(t *testing.T) {
	f := workerTestFixtures(t, 1*time.Second)
	defer f.done()

	// Assign the worker.
	require.NoError(t, f.db.WorkerSetClusters(f.ctx, f.worker, []string{f.cluster.UUID}))

	// Delete the cluster.
	require.NoError(t, f.db.DeleteWorkerCluster(f.ctx, f.cluster.UUID))

	// Check the Worker has been unassigned from the cluster.
	w, err := f.db.FetchWorker(f.ctx, f.worker.UUID)
	require.NoError(t, err)
	assert.Empty(t, w.Clusters)
}
