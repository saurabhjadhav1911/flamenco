package worker

// SPDX-License-Identifier: GPL-3.0-or-later

import (
	"context"

	"projects.blender.org/studio/flamenco/pkg/api"
)

func (w *Worker) gotoStateRestart(ctx context.Context) {
	w.stateMutex.Lock()
	defer w.stateMutex.Unlock()

	w.state = api.WorkerStatusRestart
	w.requestShutdown(true)
}
