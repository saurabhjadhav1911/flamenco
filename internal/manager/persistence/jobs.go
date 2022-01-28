package persistence

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
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"gitlab.com/blender/flamenco-ng-poc/internal/manager/job_compilers"
	"gitlab.com/blender/flamenco-ng-poc/pkg/api"
	"gorm.io/gorm"
)

type Job struct {
	gorm.Model
	UUID string `gorm:"type:char(36);not null;unique;index"`

	Name     string `gorm:"type:varchar(64);not null"`
	JobType  string `gorm:"type:varchar(32);not null"`
	Priority int8   `gorm:"type:smallint;not null"`
	Status   string `gorm:"type:varchar(32);not null"` // See JobStatusXxxx consts in openapi_types.gen.go

	Settings StringInterfaceMap `gorm:"type:jsonb"`
	Metadata StringStringMap    `gorm:"type:jsonb"`
}

type StringInterfaceMap map[string]interface{}
type StringStringMap map[string]string

type Task struct {
	gorm.Model

	Name     string `gorm:"type:varchar(64);not null"`
	Type     string `gorm:"type:varchar(32);not null"`
	JobID    uint   `gorm:"not null"`
	Job      *Job   `gorm:"foreignkey:JobID;references:ID;constraint:OnDelete:CASCADE;not null"`
	Priority int    `gorm:"type:smallint;not null"`
	Status   string `gorm:"type:varchar(16);not null"`

	// TODO: include info about which worker is/was working on this.

	// Dependencies are tasks that need to be completed before this one can run.
	Dependencies []*Task `gorm:"many2many:task_dependencies;"`

	Commands Commands `gorm:"type:jsonb"`
}

type Commands []Command

type Command struct {
	Type       string             `json:"type"`
	Parameters StringInterfaceMap `json:"parameters"`
}

func (c Commands) Value() (driver.Value, error) {
	return json.Marshal(c)
}
func (c *Commands) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &c)
}

func (js StringInterfaceMap) Value() (driver.Value, error) {
	return json.Marshal(js)
}
func (js *StringInterfaceMap) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &js)
}

func (js StringStringMap) Value() (driver.Value, error) {
	return json.Marshal(js)
}
func (js *StringStringMap) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &js)
}

func (db *DB) StoreJob(ctx context.Context, authoredJob job_compilers.AuthoredJob) error {
	return db.gormDB.Transaction(func(tx *gorm.DB) error {
		// TODO: separate conversion of struct types from storing things in the database.
		dbJob := Job{
			UUID:     authoredJob.JobID,
			Name:     authoredJob.Name,
			JobType:  authoredJob.JobType,
			Priority: int8(authoredJob.Priority),
			Settings: StringInterfaceMap(authoredJob.Settings),
			Metadata: StringStringMap(authoredJob.Metadata),
		}

		if err := db.gormDB.Create(&dbJob).Error; err != nil {
			return fmt.Errorf("error storing job: %v", err)
		}

		for _, authoredTask := range authoredJob.Tasks {
			var commands []Command
			for _, authoredCommand := range authoredTask.Commands {
				commands = append(commands, Command{
					Type:       authoredCommand.Type,
					Parameters: StringInterfaceMap(authoredCommand.Parameters),
				})
			}

			dbTask := Task{
				Name:     authoredTask.Name,
				Type:     authoredTask.Type,
				Job:      &dbJob,
				Priority: authoredTask.Priority,
				Status:   string(api.TaskStatusProcessing), // TODO: is this the right place to set the default status?
				// TODO: store dependencies
				Commands: commands,
			}
			if err := db.gormDB.Create(&dbTask).Error; err != nil {
				return fmt.Errorf("error storing task: %v", err)
			}
		}

		return nil
	})
}

func (db *DB) FetchJob(ctx context.Context, jobID string) (*api.Job, error) {
	dbJob := Job{}
	findResult := db.gormDB.First(&dbJob, "uuid = ?", jobID)
	if findResult.Error != nil {
		return nil, findResult.Error
	}

	apiJob := api.Job{
		SubmittedJob: api.SubmittedJob{
			Name:     dbJob.Name,
			Priority: int(dbJob.Priority),
			Type:     dbJob.JobType,
		},

		Id:      dbJob.UUID,
		Created: dbJob.CreatedAt,
		Updated: dbJob.UpdatedAt,
		Status:  api.JobStatus(dbJob.Status),
	}

	apiJob.Settings = &api.JobSettings{AdditionalProperties: dbJob.Settings}
	apiJob.Metadata = &api.JobMetadata{AdditionalProperties: dbJob.Metadata}

	return &apiJob, nil
}