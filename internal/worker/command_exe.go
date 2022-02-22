package worker

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
	"context"
	"fmt"
	"os/exec"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gitlab.com/blender/flamenco-ng-poc/pkg/api"
)

// Generate mock implementation of this interface.
//go:generate go run github.com/golang/mock/mockgen -destination mocks/command_listener.gen.go -package mocks gitlab.com/blender/flamenco-ng-poc/internal/worker CommandListener

// CommandListener sends the result of commands (log, output files) to the Manager.
type CommandListener interface {
	// LogProduced sends any logging to whatever service for storing logging.
	LogProduced(ctx context.Context, taskID string, logLines ...string) error
	// OutputProduced tells the Manager there has been some output (most commonly a rendered frame or video).
	OutputProduced(ctx context.Context, taskID string, outputLocation string) error
}

type CommandExecutor struct {
	cli         CommandLineRunner
	listener    CommandListener
	timeService TimeService

	// registry maps a command name to a function that runs that command.
	registry map[string]commandCallable
}

var _ CommandRunner = (*CommandExecutor)(nil)

type commandCallable func(ctx context.Context, logger zerolog.Logger, taskID string, cmd api.Command) error

// TimeService is a service that operates on time.
type TimeService interface {
	After(duration time.Duration) <-chan time.Time
}

//go:generate go run github.com/golang/mock/mockgen -destination mocks/cli_runner.gen.go -package mocks gitlab.com/blender/flamenco-ng-poc/internal/worker CommandLineRunner
// CommandLineRunner is an interface around exec.CommandContext().
type CommandLineRunner interface {
	CommandContext(ctx context.Context, name string, arg ...string) *exec.Cmd
}

func NewCommandExecutor(cli CommandLineRunner, listener CommandListener, timeService TimeService) *CommandExecutor {
	ce := &CommandExecutor{
		cli:         cli,
		listener:    listener,
		timeService: timeService,
	}

	// Registry of supported commands. Having this as a map (instead of a big
	// switch statement) makes it possible to do things like reporting the list of
	// supported commands.
	ce.registry = map[string]commandCallable{
		"echo":           ce.cmdEcho,
		"sleep":          ce.cmdSleep,
		"blender-render": ce.cmdBlenderRender,
	}

	return ce
}

func (ce *CommandExecutor) Run(ctx context.Context, taskID string, cmd api.Command) error {
	logger := log.With().Str("task", string(taskID)).Str("command", cmd.Name).Logger()
	logger.Info().Interface("parameters", cmd.Parameters).Msg("running command")

	runner, ok := ce.registry[cmd.Name]
	if !ok {
		return fmt.Errorf("unknown command: %q", cmd.Name)
	}

	return runner(ctx, logger, taskID, cmd)
}

// cmdSettingAsStrings converts an array setting ([]interface{}) to a []string slice.
func cmdSettingAsStrings(cmd api.Command, key string) ([]string, bool) {
	setting, found := cmd.Parameters[key]
	if !found {
		return []string{}, false
	}

	if asStrSlice, ok := setting.([]string); ok {
		return asStrSlice, true
	}

	interfSlice, ok := setting.([]interface{})
	if !ok {
		return []string{}, false
	}

	strSlice := make([]string, len(interfSlice))
	for idx := range interfSlice {
		switch v := interfSlice[idx].(type) {
		case string:
			strSlice[idx] = v
		case fmt.Stringer:
			strSlice[idx] = v.String()
		default:
			strSlice[idx] = fmt.Sprintf("%v", v)
		}
	}
	return strSlice, true
}
