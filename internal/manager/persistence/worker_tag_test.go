package persistence

// SPDX-License-Identifier: GPL-3.0-or-later

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"projects.blender.org/studio/flamenco/internal/uuid"
)

func TestCreateFetchTag(t *testing.T) {
	f := workerTestFixtures(t, 1*time.Second)
	defer f.done()

	// Test fetching non-existent tag
	fetchedTag, err := f.db.FetchWorkerTag(f.ctx, "7ee21bc8-ff1a-42d2-a6b6-cc4b529b189f")
	assert.ErrorIs(t, err, ErrWorkerTagNotFound)
	assert.Nil(t, fetchedTag)

	// New tag creation is already done in the workerTestFixtures() call.
	assert.NotNil(t, f.tag)

	fetchedTag, err = f.db.FetchWorkerTag(f.ctx, f.tag.UUID)
	require.NoError(t, err)
	assert.NotNil(t, fetchedTag)

	// Test contents of fetched tag.
	assert.Equal(t, f.tag.UUID, fetchedTag.UUID)
	assert.Equal(t, f.tag.Name, fetchedTag.Name)
	assert.Equal(t, f.tag.Description, fetchedTag.Description)
	assert.Zero(t, fetchedTag.Workers)
}

func TestFetchDeleteTags(t *testing.T) {
	f := workerTestFixtures(t, 1*time.Second)
	defer f.done()

	// Single tag was created by fixture.
	has, err := f.db.HasWorkerTags(f.ctx)
	require.NoError(t, err)
	assert.True(t, has, "expecting HasWorkerTags to return true")

	secondTag := WorkerTag{
		UUID:        uuid.New(),
		Name:        "arbeiderstag",
		Description: "Worker tag in Dutch",
	}

	require.NoError(t, f.db.CreateWorkerTag(f.ctx, &secondTag))

	allTags, err := f.db.FetchWorkerTags(f.ctx)
	require.NoError(t, err)

	require.Len(t, allTags, 2)
	var allTagIDs [2]string
	for idx := range allTags {
		allTagIDs[idx] = allTags[idx].UUID
	}
	assert.Contains(t, allTagIDs, f.tag.UUID)
	assert.Contains(t, allTagIDs, secondTag.UUID)

	has, err = f.db.HasWorkerTags(f.ctx)
	require.NoError(t, err)
	assert.True(t, has, "expecting HasWorkerTags to return true")

	// Test deleting the 2nd tag.
	require.NoError(t, f.db.DeleteWorkerTag(f.ctx, secondTag.UUID))

	allTags, err = f.db.FetchWorkerTags(f.ctx)
	require.NoError(t, err)
	require.Len(t, allTags, 1)
	assert.Equal(t, f.tag.UUID, allTags[0].UUID)

	// Test deleting the 1st tag.
	require.NoError(t, f.db.DeleteWorkerTag(f.ctx, f.tag.UUID))
	has, err = f.db.HasWorkerTags(f.ctx)
	require.NoError(t, err)
	assert.False(t, has, "expecting HasWorkerTags to return false")
}

func TestAssignUnassignWorkerTags(t *testing.T) {
	f := workerTestFixtures(t, 1*time.Second)
	defer f.done()

	assertTags := func(msgLabel string, tagUUIDs ...string) {
		w, err := f.db.FetchWorker(f.ctx, f.worker.UUID)
		require.NoError(t, err)

		// Catch doubly-reported tags, as the maps below would hide those cases.
		assert.Len(t, w.Tags, len(tagUUIDs), msgLabel)

		expectTags := make(map[string]bool)
		for _, cid := range tagUUIDs {
			expectTags[cid] = true
		}

		actualTags := make(map[string]bool)
		for _, c := range w.Tags {
			actualTags[c.UUID] = true
		}

		assert.Equal(t, expectTags, actualTags, msgLabel)
	}

	secondTag := WorkerTag{
		UUID:        uuid.New(),
		Name:        "arbeiderstag",
		Description: "Worker tag in Dutch",
	}

	require.NoError(t, f.db.CreateWorkerTag(f.ctx, &secondTag))

	// By default the Worker should not be part of a tag.
	assertTags("default tag assignment")

	require.NoError(t, f.db.WorkerSetTags(f.ctx, f.worker, []string{f.tag.UUID}))
	assertTags("setting one tag", f.tag.UUID)

	// Double assignments should also just work.
	require.NoError(t, f.db.WorkerSetTags(f.ctx, f.worker, []string{f.tag.UUID, f.tag.UUID}))
	assertTags("setting twice the same tag", f.tag.UUID)

	// Multiple tag memberships.
	require.NoError(t, f.db.WorkerSetTags(f.ctx, f.worker, []string{f.tag.UUID, secondTag.UUID}))
	assertTags("setting two different tags", f.tag.UUID, secondTag.UUID)

	// Remove memberships.
	require.NoError(t, f.db.WorkerSetTags(f.ctx, f.worker, []string{secondTag.UUID}))
	assertTags("unassigning from first tag", secondTag.UUID)
	require.NoError(t, f.db.WorkerSetTags(f.ctx, f.worker, []string{}))
	assertTags("unassigning from second tag")
}

func TestSaveWorkerTag(t *testing.T) {
	f := workerTestFixtures(t, 1*time.Second)
	defer f.done()

	f.tag.Name = "übertag"
	f.tag.Description = "ʻO kēlā hui ma laila"
	require.NoError(t, f.db.SaveWorkerTag(f.ctx, f.tag))

	fetched, err := f.db.FetchWorkerTag(f.ctx, f.tag.UUID)
	require.NoError(t, err)
	assert.Equal(t, f.tag.Name, fetched.Name)
	assert.Equal(t, f.tag.Description, fetched.Description)
}

func TestDeleteWorkerTagWithWorkersAssigned(t *testing.T) {
	f := workerTestFixtures(t, 1*time.Second)
	defer f.done()

	// Assign the worker.
	require.NoError(t, f.db.WorkerSetTags(f.ctx, f.worker, []string{f.tag.UUID}))

	// Delete the tag.
	require.NoError(t, f.db.DeleteWorkerTag(f.ctx, f.tag.UUID))

	// Check the Worker has been unassigned from the tag.
	w, err := f.db.FetchWorker(f.ctx, f.worker.UUID)
	require.NoError(t, err)
	assert.Empty(t, w.Tags)
}
