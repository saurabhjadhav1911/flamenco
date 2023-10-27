package worker

// SPDX-License-Identifier: GPL-3.0-or-later

/* This file contains the "cli" command in the "misc" type group. */

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/rs/zerolog"

	"projects.blender.org/studio/flamenco/pkg/api"
)

type CliParams struct {
	exe  string   // Executable to run.
	args []string // Arguments for the executable.
}

// cmdCLI runs an arbitrary executable with arguments.
func (ce *CommandExecutor) cmdCLI(ctx context.Context, logger zerolog.Logger, taskID string, cmd api.Command) error {
	cmdCtx, cmdCtxCancel := context.WithCancel(ctx)
	defer cmdCtxCancel()

	execCmd, err := ce.cmdCLICommand(cmdCtx, logger, taskID, cmd)
	if err != nil {
		return err
	}

	logChunker := NewLogChunker(taskID, ce.listener, ce.timeService)
	subprocessErr := ce.cli.RunWithTextOutput(ctx, logger, execCmd, logChunker, nil)

	if subprocessErr != nil {
		logger.Error().Err(subprocessErr).
			Int("exitCode", execCmd.ProcessState.ExitCode()).
			Msg("command exited abnormally")
		return subprocessErr
	}

	logger.Info().Msg("command exited succesfully")
	return nil
}

func (ce *CommandExecutor) cmdCLICommand(
	ctx context.Context,
	logger zerolog.Logger,
	taskID string,
	cmd api.Command,
) (*exec.Cmd, error) {
	parameters, err := cmdCLIParams(logger, cmd)
	if err != nil {
		return nil, err
	}

	execCmd := ce.cli.CommandContext(ctx, parameters.exe, parameters.args...)
	if execCmd == nil {
		logger.Error().Msg("unable to create command executor")
		return nil, ErrNoExecCmd
	}
	logger.Info().
		Str("execCmd", execCmd.String()).
		Msg("going to execute CLI command")

	if err := ce.listener.LogProduced(ctx, taskID, fmt.Sprintf("going to run: %s %q", parameters.exe, parameters.args)); err != nil {
		return nil, err
	}

	return execCmd, nil
}

func cmdCLIParams(logger zerolog.Logger, cmd api.Command) (CliParams, error) {
	var (
		parameters CliParams
		ok         bool
	)

	if parameters.exe, ok = cmdParameter[string](cmd, "exe"); !ok || parameters.exe == "" {
		logger.Warn().Interface("command", cmd).Msg("missing 'exe' parameter")
		return parameters, NewParameterMissingError("exe", cmd)
	}
	if parameters.args, ok = cmdParameterAsStrings(cmd, "args"); !ok {
		logger.Warn().Interface("command", cmd).Msg("invalid 'args' parameter")
		return parameters, NewParameterInvalidError("args", cmd, "cannot convert to list of strings")
	}

	return parameters, nil
}
