package worker

// SPDX-License-Identifier: GPL-3.0-or-later

import (
	"testing"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/stretchr/testify/assert"
)

func mockedClock(t *testing.T) *clock.Mock {
	c := clock.NewMock()
	now, err := time.ParseInLocation("2006-01-02T15:04:05", "2006-01-02T15:04:05", time.Local)
	assert.NoError(t, err)
	c.Set(now)
	return c
}
