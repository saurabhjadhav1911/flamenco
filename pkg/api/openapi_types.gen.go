// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package api

import (
	"time"
)

const (
	Worker_authScopes = "worker_auth.Scopes"
)

// Defines values for AvailableJobSettingSubtype.
const (
	AvailableJobSettingSubtypeDirPath AvailableJobSettingSubtype = "dir_path"

	AvailableJobSettingSubtypeFileName AvailableJobSettingSubtype = "file_name"

	AvailableJobSettingSubtypeFilePath AvailableJobSettingSubtype = "file_path"

	AvailableJobSettingSubtypeHashedFilePath AvailableJobSettingSubtype = "hashed_file_path"
)

// Defines values for AvailableJobSettingType.
const (
	AvailableJobSettingTypeBool AvailableJobSettingType = "bool"

	AvailableJobSettingTypeFloat AvailableJobSettingType = "float"

	AvailableJobSettingTypeInt32 AvailableJobSettingType = "int32"

	AvailableJobSettingTypeString AvailableJobSettingType = "string"
)

// Defines values for JobStatus.
const (
	JobStatusActive JobStatus = "active"

	JobStatusArchived JobStatus = "archived"

	JobStatusArchiving JobStatus = "archiving"

	JobStatusCancelRequested JobStatus = "cancel-requested"

	JobStatusCanceled JobStatus = "canceled"

	JobStatusCompleted JobStatus = "completed"

	JobStatusConstructionFailed JobStatus = "construction-failed"

	JobStatusFailRequested JobStatus = "fail-requested"

	JobStatusFailed JobStatus = "failed"

	JobStatusPaused JobStatus = "paused"

	JobStatusQueued JobStatus = "queued"

	JobStatusRequeued JobStatus = "requeued"

	JobStatusUnderConstruction JobStatus = "under-construction"

	JobStatusWaitingForFiles JobStatus = "waiting-for-files"
)

// Defines values for TaskStatus.
const (
	TaskStatusActive TaskStatus = "active"

	TaskStatusCancelRequested TaskStatus = "cancel-requested"

	TaskStatusCanceled TaskStatus = "canceled"

	TaskStatusCompleted TaskStatus = "completed"

	TaskStatusFailRequested TaskStatus = "fail-requested"

	TaskStatusFailed TaskStatus = "failed"

	TaskStatusPaused TaskStatus = "paused"

	TaskStatusProcessing TaskStatus = "processing"

	TaskStatusQueued TaskStatus = "queued"

	TaskStatusSoftFailed TaskStatus = "soft-failed"
)

// AssignedTask is a task as it is received by the Worker.
type AssignedTask struct {
	Commands    []Command  `json:"commands"`
	Id          string     `json:"id"`
	Job         string     `json:"job"`
	JobPriority int        `json:"job_priority"`
	JobType     string     `json:"job_type"`
	Name        string     `json:"name"`
	Priority    int        `json:"priority"`
	Status      TaskStatus `json:"status"`
	TaskType    string     `json:"task_type"`
	User        string     `json:"user"`
}

// Single setting of a Job types.
type AvailableJobSetting struct {
	// When given, limit the valid values to these choices. Only usable with string type.
	Choices *[]string `json:"choices,omitempty"`

	// The default value shown to the user when determining this setting.
	Default *interface{} `json:"default,omitempty"`

	// Whether to allow editing this setting after the job has been submitted. Would imply deleting all existing tasks for this job, and recompiling it.
	Editable *bool `json:"editable,omitempty"`

	// Identifier for the setting, must be unique within the job type.
	Key string `json:"key"`

	// Whether to immediately reject a job definition, of this type, without this particular setting.
	Required *bool `json:"required,omitempty"`

	// Sub-type of the job setting. Currently only available for string types. `HASHED_FILE_PATH` is a directory path + `"/######"` appended.
	Subtype *AvailableJobSettingSubtype `json:"subtype,omitempty"`

	// Type of job setting, must be usable as IDProperty type in Blender. No nested structures (arrays, dictionaries) are supported.
	Type AvailableJobSettingType `json:"type"`

	// Whether to show this setting in the UI of a job submitter (like a Blender add-on). Set to `false` when it is an internal setting that shouldn't be shown to end users.
	Visible *bool `json:"visible,omitempty"`
}

// Sub-type of the job setting. Currently only available for string types. `HASHED_FILE_PATH` is a directory path + `"/######"` appended.
type AvailableJobSettingSubtype string

// Type of job setting, must be usable as IDProperty type in Blender. No nested structures (arrays, dictionaries) are supported.
type AvailableJobSettingType string

// Job type supported by this Manager, and its parameters.
type AvailableJobType struct {
	Label    string                `json:"label"`
	Name     string                `json:"name"`
	Settings []AvailableJobSetting `json:"settings"`
}

// List of job types supported by this Manager.
type AvailableJobTypes struct {
	JobTypes []AvailableJobType `json:"job_types"`
}

// Command represents a single command to execute by the Worker.
type Command struct {
	Name     string                 `json:"name"`
	Settings map[string]interface{} `json:"settings"`
}

// Error defines model for Error.
type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// Job defines model for Job.
type Job struct {
	// Embedded struct due to allOf(#/components/schemas/SubmittedJob)
	SubmittedJob `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	// Creation timestamp
	Created time.Time `json:"created"`

	// UUID of the Job
	Id string `json:"id"`

	// Creation timestamp
	Updated time.Time `json:"updated"`
}

// JobStatus defines model for JobStatus.
type JobStatus string

// RegisteredWorker defines model for RegisteredWorker.
type RegisteredWorker struct {
	Address            string   `json:"address"`
	Id                 string   `json:"id"`
	LastActivity       string   `json:"last_activity"`
	Nickname           string   `json:"nickname"`
	Platform           string   `json:"platform"`
	Software           string   `json:"software"`
	Status             string   `json:"status"`
	SupportedTaskTypes []string `json:"supported_task_types"`
}

// SecurityError defines model for SecurityError.
type SecurityError struct {
	Message string `json:"message"`
}

// Job definition submitted to Flamenco.
type SubmittedJob struct {
	Metadata *map[string]interface{} `json:"metadata,omitempty"`
	Name     string                  `json:"name"`
	Priority int                     `json:"priority"`
	Settings *map[string]interface{} `json:"settings,omitempty"`
	Status   *JobStatus              `json:"status,omitempty"`
	Type     string                  `json:"type"`
}

// TaskStatus defines model for TaskStatus.
type TaskStatus string

// WorkerRegistration defines model for WorkerRegistration.
type WorkerRegistration struct {
	Nickname           string   `json:"nickname"`
	Platform           string   `json:"platform"`
	Secret             string   `json:"secret"`
	SupportedTaskTypes []string `json:"supported_task_types"`
}

// SubmitJobJSONBody defines parameters for SubmitJob.
type SubmitJobJSONBody SubmittedJob

// RegisterWorkerJSONBody defines parameters for RegisterWorker.
type RegisterWorkerJSONBody WorkerRegistration

// SubmitJobJSONRequestBody defines body for SubmitJob for application/json ContentType.
type SubmitJobJSONRequestBody SubmitJobJSONBody

// RegisterWorkerJSONRequestBody defines body for RegisterWorker for application/json ContentType.
type RegisterWorkerJSONRequestBody RegisterWorkerJSONBody
