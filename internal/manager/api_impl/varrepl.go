package api_impl

// SPDX-License-Identifier: GPL-3.0-or-later

import (
	"sync"

	"git.blender.org/flamenco/internal/manager/config"
	"git.blender.org/flamenco/internal/manager/persistence"
	"git.blender.org/flamenco/pkg/api"
)

//go:generate go run github.com/golang/mock/mockgen -destination mocks/varrepl.gen.go -package mocks git.blender.org/flamenco/internal/manager/api_impl VariableReplacer
type VariableReplacer interface {
	ExpandVariables(inputChannel <-chan string, outputChannel chan<- string, audience config.VariableAudience, platform config.VariablePlatform)
	ResolveVariables(audience config.VariableAudience, platform config.VariablePlatform) map[string]config.ResolvedVariable
	NewVariableToValueConverter(audience config.VariableAudience, platform config.VariablePlatform) *config.ValueToVariableReplacer
}

// replaceTaskVariables performs variable replacement for worker tasks.
func replaceTaskVariables(replacer VariableReplacer, task api.AssignedTask, worker persistence.Worker) api.AssignedTask {
	feeder := make(chan string, 1)
	receiver := make(chan string, 1)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		replacer.ExpandVariables(feeder, receiver,
			config.VariableAudienceWorkers, config.VariablePlatform(worker.Platform))
	}()

	for cmdIndex, cmd := range task.Commands {
		for key, value := range cmd.Parameters {
			switch v := value.(type) {
			case string:
				feeder <- v
				task.Commands[cmdIndex].Parameters[key] = <-receiver

			case []string:
				replaced := make([]string, len(v))
				for idx := range v {
					feeder <- v[idx]
					replaced[idx] = <-receiver
				}
				task.Commands[cmdIndex].Parameters[key] = replaced

			case []interface{}:
				replaced := make([]interface{}, len(v))
				for idx := range v {
					switch itemValue := v[idx].(type) {
					case string:
						feeder <- itemValue
						replaced[idx] = <-receiver
					default:
						replaced[idx] = itemValue
					}
				}
				task.Commands[cmdIndex].Parameters[key] = replaced

			default:
				continue
			}
		}
	}

	close(feeder)
	wg.Wait()
	close(receiver)

	return task
}

// replaceTwoWayVariables replaces values with their variables.
// For example, when there is a variable `render = /render/flamenco`, an output
// path `/render/flamenco/output.png` will be replaced with
// `{render}/output.png`
//
// NOTE: this updates the job in place.
func replaceTwoWayVariables(replacer VariableReplacer, job *api.SubmittedJob) {
	valueToVariable := replacer.NewVariableToValueConverter(
		config.VariableAudienceWorkers,
		config.VariablePlatform(job.SubmitterPlatform),
	)

	// Only replace variables in settings and metadata, not in other job fields.
	if job.Settings != nil {
		for settingKey, settingValue := range job.Settings.AdditionalProperties {
			stringValue, ok := settingValue.(string)
			if !ok {
				continue
			}
			newValue := valueToVariable.Replace(stringValue)
			job.Settings.AdditionalProperties[settingKey] = newValue
		}
	}
	if job.Metadata != nil {
		for metaKey, metaValue := range job.Metadata.AdditionalProperties {
			newValue := valueToVariable.Replace(metaValue)
			job.Metadata.AdditionalProperties[metaKey] = newValue
		}
	}
}
