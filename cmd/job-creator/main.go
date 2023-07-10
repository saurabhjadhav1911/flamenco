package main

// SPDX-License-Identifier: GPL-3.0-or-later

import (
	"context"
	"errors"
	"flag"
	"io/fs"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/mattn/go-colorable"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"git.blender.org/flamenco/internal/appinfo"
	"git.blender.org/flamenco/internal/manager/config"
	"git.blender.org/flamenco/internal/manager/job_compilers"
	"git.blender.org/flamenco/internal/manager/persistence"
	"git.blender.org/flamenco/pkg/api"
)

var cliArgs struct {
	version bool
	jobUUID string
}

func main() {
	output := zerolog.ConsoleWriter{Out: colorable.NewColorableStdout(), TimeFormat: time.RFC3339}
	log.Logger = log.Output(output)
	log.Info().
		Str("version", appinfo.ApplicationVersion).
		Str("git", appinfo.ApplicationGitHash).
		Str("releaseCycle", appinfo.ReleaseCycle).
		Str("os", runtime.GOOS).
		Str("arch", runtime.GOARCH).
		Msgf("starting %v job compiler", appinfo.ApplicationName)

	parseCliArgs()
	if cliArgs.version {
		return
	}

	if cliArgs.jobUUID == "" {
		log.Fatal().Msg("give me a job UUID to regenerate tasks for")
	}

	// Load configuration.
	configService := config.NewService()
	err := configService.Load()
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		log.Error().Err(err).Msg("loading configuration")
	}

	isFirstRun, err := configService.IsFirstRun()
	switch {
	case err != nil:
		log.Fatal().Err(err).Msg("unable to determine whether this is the first run of Flamenco or not")
	case isFirstRun:
		log.Info().Msg("This seems to be your first run of Flamenco, this tool won't work.")
		return
	}

	// Construct the services.
	persist := openDB(*configService)
	defer persist.Close()

	timeService := clock.New()
	compiler, err := job_compilers.Load(timeService)
	if err != nil {
		log.Fatal().Err(err).Msg("error loading job compilers")
	}

	// The main context determines the lifetime of the application. All
	// long-running goroutines need to keep an eye on this, and stop their work
	// once it closes.
	mainCtx, mainCtxCancel := context.WithCancel(context.Background())
	defer mainCtxCancel()

	installSignalHandler(mainCtxCancel)

	recompile(mainCtx, cliArgs.jobUUID, persist, compiler)
}

// recompile regenerates the job's tasks.
func recompile(ctx context.Context, jobUUID string, db *persistence.DB, compiler *job_compilers.Service) {
	dbJob, err := db.FetchJob(ctx, jobUUID)
	if err != nil {
		log.Fatal().Err(err).Msg("could not get job from database")
	}
	logger := log.With().Str("job", jobUUID).Logger()
	logger.Info().Msg("found job")

	dbTasks, err := db.FetchTasksOfJob(ctx, dbJob)
	if err != nil {
		log.Fatal().Err(err).Msg("could not query database for tasks")
	}
	if len(dbTasks) > 0 {
		// This tool has only been tested with jobs that have had their tasks completely lost.
		log.Fatal().
			Int("numTasks", len(dbTasks)).
			Msg("this job still has tasks, this is not a situation this tool should be used in")
	}

	// Recompile the job.
	fakeSubmittedJob := constructSubmittedJob(dbJob)
	authoredJob, err := compiler.Compile(ctx, fakeSubmittedJob)
	if err != nil {
		logger.Fatal().Err(err).Msg("could not recompile job")
	}
	sanityCheck(logger, dbJob, authoredJob)

	// Store the recompiled tasks.
	if err := db.StoreAuthoredJobTaks(ctx, dbJob, authoredJob); err != nil {
		logger.Fatal().Err(err).Msg("error storing recompiled tasks")
	}
	logger.Info().Msg("new tasks have been stored")

	updateTaskStatuses(ctx, logger, db, dbJob)

	logger.Info().Msg("job recompilation seems to have worked out")
}

func constructSubmittedJob(dbJob *persistence.Job) api.SubmittedJob {
	fakeSubmittedJob := api.SubmittedJob{
		Name:              dbJob.Name,
		Priority:          dbJob.Priority,
		SubmitterPlatform: "reconstrutor", // The platform shouldn't matter, as all paths have already been replaced.
		Type:              dbJob.JobType,
		TypeEtag:          nil,

		Settings: &api.JobSettings{AdditionalProperties: make(map[string]interface{})},
		Metadata: &api.JobMetadata{AdditionalProperties: make(map[string]string)},
	}

	for key, value := range dbJob.Settings {
		fakeSubmittedJob.Settings.AdditionalProperties[key] = value
	}
	for key, value := range dbJob.Metadata {
		fakeSubmittedJob.Metadata.AdditionalProperties[key] = value
	}
	if dbJob.WorkerTag != nil {
		fakeSubmittedJob.WorkerTag = &dbJob.WorkerTag.UUID
	} else if dbJob.WorkerTagID != nil {
		panic("WorkerTagID is set, but WorkerTag is not")
	}

	return fakeSubmittedJob
}

// Check that the authored job is consistent with the original job.
func sanityCheck(logger zerolog.Logger, expect *persistence.Job, actual *job_compilers.AuthoredJob) {
	if actual.Name != expect.Name {
		logger.Fatal().
			Str("expected", expect.Name).
			Str("actual", actual.Name).
			Msg("recompilation did not produce expected name")
	}
	if actual.JobType != expect.JobType {
		logger.Fatal().
			Str("expected", expect.JobType).
			Str("actual", actual.JobType).
			Msg("recompilation did not produce expected job type")
	}
}

func updateTaskStatuses(ctx context.Context, logger zerolog.Logger, db *persistence.DB, dbJob *persistence.Job) {
	logger = logger.With().Str("jobStatus", string(dbJob.Status)).Logger()

	// Update the task statuses based on the job status. This is NOT using the
	// state machine, as these tasks are not actually going from one state to the
	// other. They are just being updated in the database.
	taskStatusMap := map[api.JobStatus]api.TaskStatus{
		api.JobStatusActive:            api.TaskStatusQueued,
		api.JobStatusCancelRequested:   api.TaskStatusCanceled,
		api.JobStatusCanceled:          api.TaskStatusCanceled,
		api.JobStatusCompleted:         api.TaskStatusCompleted,
		api.JobStatusFailed:            api.TaskStatusCanceled,
		api.JobStatusPaused:            api.TaskStatusPaused,
		api.JobStatusQueued:            api.TaskStatusQueued,
		api.JobStatusRequeueing:        api.TaskStatusQueued,
		api.JobStatusUnderConstruction: api.TaskStatusQueued,
	}
	newTaskStatus, ok := taskStatusMap[dbJob.Status]
	if !ok {
		logger.Warn().Msg("unknown job status, not touching task statuses")
		return
	}

	logger = logger.With().Str("taskStatus", string(newTaskStatus)).Logger()

	err := db.UpdateJobsTaskStatuses(ctx, dbJob, newTaskStatus, "reset task status after job reconstruction")
	if err != nil {
		logger.Fatal().Msg("could not update task statuses")
	}

	logger.Info().Msg("task statuses have been updated based on the job status")
}

func parseCliArgs() {
	var quiet, debug, trace bool

	flag.BoolVar(&cliArgs.version, "version", false, "Shows the application version, then exits.")
	flag.BoolVar(&quiet, "quiet", false, "Only log warning-level and worse.")
	flag.BoolVar(&debug, "debug", false, "Enable debug-level logging.")
	flag.BoolVar(&trace, "trace", false, "Enable trace-level logging.")
	flag.StringVar(&cliArgs.jobUUID, "job", "", "Job UUID to regenerate")

	flag.Parse()

	var logLevel zerolog.Level
	switch {
	case trace:
		logLevel = zerolog.TraceLevel
	case debug:
		logLevel = zerolog.DebugLevel
	case quiet:
		logLevel = zerolog.WarnLevel
	default:
		logLevel = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(logLevel)
}

// openDB opens the database or dies.
func openDB(configService config.Service) *persistence.DB {
	dsn := configService.Get().DatabaseDSN
	if dsn == "" {
		log.Fatal().Msg("configure the database in flamenco-manager.yaml")
	}

	dbCtx, dbCtxCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer dbCtxCancel()
	persist, err := persistence.OpenDB(dbCtx, dsn)
	if err != nil {
		log.Fatal().
			Err(err).
			Str("dsn", dsn).
			Msg("error opening database")
	}

	return persist
}

// installSignalHandler spawns a goroutine that handles incoming POSIX signals.
func installSignalHandler(cancelFunc context.CancelFunc) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	signal.Notify(signals, syscall.SIGTERM)
	go func() {
		for signum := range signals {
			log.Info().Str("signal", signum.String()).Msg("signal received, shutting down")
			cancelFunc()
		}
	}()
}
