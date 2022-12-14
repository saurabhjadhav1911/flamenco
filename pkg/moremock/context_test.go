package moremock

// SPDX-License-Identifier: GPL-3.0-or-later

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCtxMatcher(t *testing.T) {
	m := ctxWithDeadlineMatcher{}

	// Non-context types -> No match.
	assert.False(t, m.Matches(nil))
	assert.False(t, m.Matches("something else"))
	var nilValuedInterface context.Context
	assert.False(t, m.Matches(nilValuedInterface))

	// Context without deadlines -> No match.
	assert.False(t, m.Matches(context.Background()))
	assert.False(t, m.Matches(context.TODO()))

	// Deadline in the past -> No match.
	{
		past, cancel := context.WithDeadline(context.Background(), time.Now().Add(-1*time.Second))
		defer cancel()
		assert.False(t, m.Matches(past))
	}

	// Deadline in the future -> Match.
	{
		future, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
		defer cancel()
		assert.True(t, m.Matches(future))
	}

	// Timeout in the future -> Match.
	{
		future, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		assert.True(t, m.Matches(future))
	}
}
