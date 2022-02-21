package config

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
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	yaml "gopkg.in/yaml.v2"

	"gitlab.com/blender/flamenco-ng-poc/pkg/api"
)

const (
	configFilename = "flamenco-manager.yaml"

	latestConfigVersion = 3

	// // relative to the Flamenco Server Base URL:
	// jwtPublicKeysRelativeURL = "api/flamenco/jwt/public-keys"

	defaultShamanFilestorePath = "/shared/flamenco/file-store"
	defaultJobStorage          = "/shared/flamenco/jobs"
)

var (
	// ErrMissingVariablePlatform is returned when a variable doesn't declare any valid platform for a certain value.
	ErrMissingVariablePlatform = errors.New("variable's value is missing platform declaration")
	// ErrBadDirection is returned when a direction doesn't match "oneway" or "twoway"
	ErrBadDirection = errors.New("variable's direction is invalid")

	// Valid values for the "audience" tag of a ConfV2 variable.
	validAudiences = map[VariableAudience]bool{
		VariableAudienceAll:     true,
		VariableAudienceWorkers: true,
		VariableAudienceUsers:   true,
	}

	// The default configuration, use DefaultConfig() to obtain a copy.
	defaultConfig = Conf{
		Base: Base{
			Meta: ConfMeta{Version: latestConfigVersion},

			ManagerName:  "Flamenco Manager",
			Listen:       ":8080",
			ListenHTTPS:  ":8433",
			DatabaseDSN:  "host=localhost user=flamenco password=flamenco dbname=flamenco TimeZone=Europe/Amsterdam",
			TaskLogsPath: "./task-logs",
			// DownloadTaskSleep:           10 * time.Minute,
			// DownloadTaskRecheckThrottle: 10 * time.Second,
			// TaskUpdatePushMaxInterval:   5 * time.Second,
			// TaskUpdatePushMaxCount:      3000,
			// CancelTaskFetchInterval:     10 * time.Second,
			ActiveTaskTimeoutInterval:   10 * time.Minute,
			ActiveWorkerTimeoutInterval: 1 * time.Minute,
			// FlamencoStr:                 defaultServerURL,

			// // Days are assumed to be 24 hours long. This is not exactly accurate, but should
			// // be accurate enough for this type of cleanup.
			// TaskCleanupMaxAge: 14 * 24 * time.Hour,
			SSDPDiscovery:  false,                                  // Only enable after SSDP discovery has been improved (avoid finding printers).
			SSDPDeviceUUID: "64ad4c21-6042-4378-9cdf-478f88b4f990", // UUID specific for Flamenco v3.

			BlacklistThreshold:         3,
			TaskFailAfterSoftFailCount: 3,

			WorkerCleanupStatus: []string{string(api.WorkerStatusOffline)},

			TestTasks: TestTasks{
				BlenderRender: BlenderRenderConfig{
					JobStorage:   "{job_storage}/test-jobs",
					RenderOutput: "{render}/test-renders",
				},
			},

			Shaman: ShamanConfig{
				Enabled:       true,
				FileStorePath: defaultShamanFilestorePath,
				GarbageCollect: ShamanGarbageCollect{
					Period:            24 * time.Hour,
					MaxAge:            31 * 24 * time.Hour,
					ExtraCheckoutDirs: []string{},
				},
			},

			// JWT: jwtauth.Config{
			// 	DownloadKeysInterval: 1 * time.Hour,
			// },
		},

		Variables: map[string]Variable{
			"blender": {
				Direction: "oneway",
				Values: VariableValues{
					VariableValue{Platform: "linux", Value: "/linux/path/to/blender --factory-startup --background"},
					VariableValue{Platform: "windows", Value: "C:/windows/path/to/blender.exe --factory-startup --background"},
					VariableValue{Platform: "darwin", Value: "/Volumes/Applications/Blender/blender --factory-startup --background"},
				},
			},
			"ffmpeg": {
				Direction: "oneway",
				Values: VariableValues{
					VariableValue{Platform: "linux", Value: "/usr/bin/ffmpeg"},
					VariableValue{Platform: "windows", Value: "C:/windows/path/to/ffmpeg.exe"},
					VariableValue{Platform: "darwin", Value: "/Volumes/Applications/FFmpeg/ffmpeg"},
				},
			},
			"job_storage": {
				Direction: "twoway",
				Values: VariableValues{
					VariableValue{Platform: "linux", Value: "/shared/flamenco/jobs"},
					VariableValue{Platform: "windows", Value: "S:/flamenco/jobs"},
					VariableValue{Platform: "darwin", Value: "/Volumes/Shared/flamenco/jobs"},
				},
			},
			"render": {
				Direction: "twoway",
				Values: VariableValues{
					VariableValue{Platform: "linux", Value: "/shared/flamenco/render"},
					VariableValue{Platform: "windows", Value: "S:/flamenco/render"},
					VariableValue{Platform: "darwin", Value: "/Volumes/Shared/flamenco/render"},
				},
			},
		},
	}
)

// BlenderRenderConfig represents the configuration required for a test render.
type BlenderRenderConfig struct {
	JobStorage   string `yaml:"job_storage"`
	RenderOutput string `yaml:"render_output"`
}

// TestTasks represents the 'test_tasks' key in the Manager's configuration file.
type TestTasks struct {
	BlenderRender BlenderRenderConfig `yaml:"test_blender_render"`
}

// ConfMeta contains configuration file metadata.
type ConfMeta struct {
	// Version of the config file structure.
	Version int `yaml:"version"`
}

// Base contains those settings that are shared by all configuration versions.
type Base struct {
	Meta ConfMeta `yaml:"_meta"`

	ManagerName  string `yaml:"manager_name"`
	DatabaseDSN  string `yaml:"database_url"`
	TaskLogsPath string `yaml:"task_logs_path"`
	Listen       string `yaml:"listen"`
	ListenHTTPS  string `yaml:"listen_https"`
	OwnURL       string `yaml:"own_url"` // sent to workers via SSDP/UPnP

	// TLS certificate management. TLSxxx has priority over ACME.
	TLSKey         string `yaml:"tlskey"`
	TLSCert        string `yaml:"tlscert"`
	ACMEDomainName string `yaml:"acme_domain_name"` // for the ACME Let's Encrypt client

	ActiveTaskTimeoutInterval   time.Duration `yaml:"active_task_timeout_interval"`
	ActiveWorkerTimeoutInterval time.Duration `yaml:"active_worker_timeout_interval"`

	WorkerCleanupMaxAge time.Duration `yaml:"worker_cleanup_max_age"`
	WorkerCleanupStatus []string      `yaml:"worker_cleanup_status"`

	/* This many failures (on a given job+task type combination) will ban a worker
	 * from that task type on that job. */
	BlacklistThreshold int `yaml:"blacklist_threshold"`

	// When this many workers have tried the task and failed, it will be hard-failed
	// (even when there are workers left that could technically retry the task).
	TaskFailAfterSoftFailCount int `yaml:"task_fail_after_softfail_count"`

	SSDPDiscovery  bool   `yaml:"ssdp_discovery"`
	SSDPDeviceUUID string `yaml:"ssdp_device_uuid"`

	TestTasks TestTasks `yaml:"test_tasks"`

	// Shaman configuration settings.
	Shaman ShamanConfig `yaml:"shaman"`

	// Authentication settings.
	// JWT                      jwtauth.Config `yaml:"user_authentication"`
	WorkerRegistrationSecret string `yaml:"worker_registration_secret"`

	// Dynamic worker pools (Azure Batch, Google Compute, AWS, that sort).
	// DynamicPoolPlatforms *dppoller.Config `yaml:"dynamic_pool_platforms,omitempty"`

	// Websetup *WebsetupConf `yaml:"websetup,omitempty"`
}

type ShamanConfig struct {
	Enabled        bool                 `yaml:"enabled"`
	FileStorePath  string               `yaml:"fileStorePath"`
	GarbageCollect ShamanGarbageCollect `yaml:"garbageCollect"`
}

// GarbageCollect contains the config options for the GC.
type ShamanGarbageCollect struct {
	// How frequently garbage collection is performed on the file store:
	Period time.Duration `yaml:"period"`
	// How old files must be before they are GC'd:
	MaxAge time.Duration `yaml:"maxAge"`
	// Paths to check for symlinks before GC'ing files.
	ExtraCheckoutDirs []string `yaml:"extraCheckoutPaths"`

	// Used by the -gc CLI arg to silently disable the garbage collector
	// while we're performing a manual sweep.
	SilentlyDisable bool `yaml:"-"`
}

// Conf is the latest version of the configuration.
// Currently it is version 3.
type Conf struct {
	Base `yaml:",inline"`

	// Variable name → Variable definition
	Variables map[string]Variable `yaml:"variables"`

	// audience + platform + variable name → variable value.
	// Used to look up variables for a given platform and audience.
	// The 'audience' is never "all" or ""; only concrete audiences are stored here.
	VariablesLookup map[VariableAudience]map[string]map[string]string `yaml:"-"`
}

// Variable defines a configuration variable.
type Variable struct {
	// Either "oneway" or "twoway"
	Direction string `yaml:"direction" json:"direction"`
	// Mapping from variable value to audience/platform definition.
	Values VariableValues `yaml:"values" json:"values"`
}

// VariableValues is the list of values of a variable.
type VariableValues []VariableValue

// VariableValue defines which audience and platform see which value.
type VariableValue struct {
	// Audience defines who will use this variable, either "all", "workers", or "users". Empty string is "all".
	Audience VariableAudience `yaml:"audience,omitempty" json:"audience,omitempty"`

	// Platforms that use this value. Only one of "Platform" and "Platforms" may be set.
	Platform  string   `yaml:"platform,omitempty" json:"platform,omitempty"`
	Platforms []string `yaml:"platforms,omitempty,flow" json:"platforms,omitempty,flow"`

	// The actual value of the variable for this audience+platform.
	Value string `yaml:"value" json:"value"`
}

// WebsetupConf are settings used by the web setup mode.
// type WebsetupConf struct {
// 	// When true, the websetup will hide certain settings that are infrastructure-specific.
// 	// For example, it hides MongoDB choice, port numbers, task log directory, all kind of
// 	// hosting-specific things. This is used, for example, by the automated Azure deployment
// 	// to avoid messing up settings that are specific to that particular installation.
// 	HideInfraSettings bool `yaml:"hide_infra_settings"`
// }

// getConf parses flamenco-manager.yaml and returns its contents as a Conf object.
func getConf() (Conf, error) {
	return loadConf(configFilename)
}

// DefaultConfig returns a copy of the default configuration.
func DefaultConfig() Conf {
	c := defaultConfig
	c.Meta.Version = latestConfigVersion
	c.constructVariableLookupTable()
	return c
}

// loadConf parses the given file and returns its contents as a Conf object.
func loadConf(filename string) (Conf, error) {
	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		return DefaultConfig(), err
	}

	// First parse attempt, find the version.
	baseConf := Base{}
	if err := yaml.Unmarshal(yamlFile, &baseConf); err != nil {
		return Conf{}, fmt.Errorf("unable to parse %s: %w", filename, err)
	}

	// Versioning was supported from Flamenco config v1 to v2, but not further.
	if baseConf.Meta.Version != latestConfigVersion {
		return Conf{}, fmt.Errorf(
			"configuration file %s version %d, but only version %d is supported",
			filename, baseConf.Meta.Version, latestConfigVersion)
	}

	// Second parse attempt, based on the version found.
	c := DefaultConfig()
	if err := yaml.Unmarshal(yamlFile, &c); err != nil {
		return c, fmt.Errorf("unable to parse %s: %w", filename, err)
	}

	c.constructVariableLookupTable()
	c.parseURLs()
	c.checkDatabase()
	c.checkVariables()
	c.checkTLS()

	return c, nil
}

func (c *Conf) constructVariableLookupTable() {
	lookup := map[VariableAudience]map[string]map[string]string{}

	// Construct a list of all audiences except "" and "all"
	concreteAudiences := []VariableAudience{}
	isWildcard := map[VariableAudience]bool{"": true, VariableAudienceAll: true}
	for audience := range validAudiences {
		if isWildcard[audience] {
			continue
		}
		concreteAudiences = append(concreteAudiences, audience)
	}
	log.Debug().
		Interface("concreteAudiences", concreteAudiences).
		Interface("isWildcard", isWildcard).
		Msg("constructing variable lookup table")

	// setValue expands wildcard audiences into concrete ones.
	var setValue func(audience VariableAudience, platform, name, value string)
	setValue = func(audience VariableAudience, platform, name, value string) {
		if isWildcard[audience] {
			for _, aud := range concreteAudiences {
				setValue(aud, platform, name, value)
			}
			return
		}

		if lookup[audience] == nil {
			lookup[audience] = map[string]map[string]string{}
		}
		if lookup[audience][platform] == nil {
			lookup[audience][platform] = map[string]string{}
		}
		log.Debug().
			Str("audience", string(audience)).
			Str("platform", platform).
			Str("name", name).
			Str("value", value).
			Msg("setting variable")
		lookup[audience][platform][name] = value
	}

	// Construct the lookup table for each audience+platform+name
	for name, variable := range c.Variables {
		log.Debug().
			Str("name", name).
			Interface("variable", variable).
			Msg("handling variable")
		for _, value := range variable.Values {

			// Two-way values should not end in path separator.
			// Given a variable 'apps' with value '/path/to/apps',
			// '/path/to/apps/blender' should be remapped to '{apps}/blender'.
			if variable.Direction == "twoway" {
				if strings.Contains(value.Value, "\\") {
					log.Warn().
						Str("variable", name).
						Str("audience", string(value.Audience)).
						Str("platform", value.Platform).
						Str("value", value.Value).
						Msg("Backslash found in variable value. Change paths to use forward slashes instead.")
				}
				value.Value = strings.TrimRight(value.Value, "/")
			}

			if value.Platform != "" {
				setValue(value.Audience, value.Platform, name, value.Value)
			}
			for _, platform := range value.Platforms {
				setValue(value.Audience, platform, name, value.Value)
			}
		}
	}
	log.Debug().
		Interface("variables", c.Variables).
		Interface("lookup", lookup).
		Msg("constructed lookup table")
	c.VariablesLookup = lookup
}

// ExpandVariables converts "{variable name}" to the value that belongs to the given audience and platform.
func (c *Conf) ExpandVariables(valueToExpand string, audience VariableAudience, platform string) string {
	audienceMap := c.VariablesLookup[audience]
	if audienceMap == nil {
		log.Warn().
			Str("valueToExpand", valueToExpand).
			Str("audience", string(audience)).
			Str("platform", platform).
			Msg("no variables defined for this audience")
		return valueToExpand
	}

	platformMap := audienceMap[platform]
	if platformMap == nil {
		log.Warn().
			Str("valueToExpand", valueToExpand).
			Str("audience", string(audience)).
			Str("platform", platform).
			Msg("no variables defined for this platform given this audience")
		return valueToExpand
	}

	// Variable replacement
	for varname, varvalue := range platformMap {
		placeholder := fmt.Sprintf("{%s}", varname)
		valueToExpand = strings.Replace(valueToExpand, placeholder, varvalue, -1)
	}

	return valueToExpand
}

// checkVariables performs some basic checks on variable definitions.
// Note that the returned error only reflects the last-found error.
// All errors are logged, though.
func (c *Conf) checkVariables() error {
	var err error

	directionNames := []string{"oneway", "twoway"}
	validDirections := map[string]bool{}
	for _, direction := range directionNames {
		validDirections[direction] = true
	}

	for name, variable := range c.Variables {
		if !validDirections[variable.Direction] {
			log.Error().
				Str("name", name).
				Str("direction", variable.Direction).
				Strs("validChoices", directionNames).
				Msg("variable has invalid direction")
			err = ErrBadDirection
		}
		for valueIndex, value := range variable.Values {
			// No platforms at all.
			if value.Platform == "" && len(value.Platforms) == 0 {
				log.Error().
					Str("name", name).
					Interface("value", value).
					Msg("variable has a platformless value")
				err = ErrMissingVariablePlatform
				continue
			}

			// Both Platform and Platforms.
			if value.Platform != "" && len(value.Platforms) > 0 {
				log.Warn().
					Str("name", name).
					Interface("value", value).
					Str("platform", value.Platform).
					Strs("platforms", value.Platforms).
					Msg("variable has a both 'platform' and 'platforms' set")
				value.Platforms = append(value.Platforms, value.Platform)
				value.Platform = ""
			}

			if value.Audience == "" {
				value.Audience = "all"
			} else if !validAudiences[value.Audience] {
				log.Error().
					Str("name", name).
					Interface("value", value).
					Str("audience", string(value.Audience)).
					Msg("variable invalid audience")
			}

			variable.Values[valueIndex] = value
		}
	}

	return err
}

func (c *Conf) checkDatabase() {
	c.DatabaseDSN = strings.TrimSpace(c.DatabaseDSN)
}

// Overwrite stores this configuration object as flamenco-manager.yaml.
func (c *Conf) Overwrite() error {
	tempFilename := configFilename + "~"
	if err := c.Write(tempFilename); err != nil {
		return fmt.Errorf("error writing config to %s: %w", tempFilename, err)
	}
	if err := os.Rename(tempFilename, configFilename); err != nil {
		return fmt.Errorf("error moving %s to %s: %w", tempFilename, configFilename, err)
	}

	log.Info().Str("filename", configFilename).Msg("saved configuration to file")
	return nil
}

// Write saves the current in-memory configuration to a YAML file.
func (c *Conf) Write(filename string) error {
	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	fmt.Fprintln(f, "# Configuration file for Flamenco Manager.")
	fmt.Fprintln(f, "# For an explanation of the fields, refer to flamenco-manager-example.yaml")
	fmt.Fprintln(f, "#")
	fmt.Fprintln(f, "# NOTE: this file will be overwritten by Flamenco Manager's web-based configuration system.")
	fmt.Fprintln(f, "#")
	now := time.Now()
	fmt.Fprintf(f, "# This file was written on %s\n\n", now.Format("2006-01-02 15:04:05 -07:00"))

	n, err := f.Write(data)
	if err != nil {
		return err
	}
	if n < len(data) {
		return io.ErrShortWrite
	}
	if err = f.Close(); err != nil {
		return err
	}

	log.Debug().Str("filename", filename).Msg("config file written")
	return nil
}

// HasCustomTLS returns true if both the TLS certificate and key files are configured.
func (c *Conf) HasCustomTLS() bool {
	return c.TLSCert != "" && c.TLSKey != ""
}

// HasTLS returns true if either a custom certificate or ACME/Let's Encrypt is used.
func (c *Conf) HasTLS() bool {
	return c.ACMEDomainName != "" || c.HasCustomTLS()
}

func (c *Conf) checkTLS() {
	hasTLS := c.HasCustomTLS()

	if hasTLS && c.ListenHTTPS == "" {
		c.ListenHTTPS = c.Listen
		c.Listen = ""
	}

	if !hasTLS || c.ACMEDomainName == "" {
		return
	}

	log.Warn().
		Str("tlscert", c.TLSCert).
		Str("tlskey", c.TLSKey).
		Str("acme_domain_name", c.ACMEDomainName).
		Msg("ACME/Let's Encrypt will not be used because custom certificate is specified")
	c.ACMEDomainName = ""
}

func (c *Conf) parseURLs() {
	// var err error
	// if jwtURL, err := c.Flamenco.Parse(jwtPublicKeysRelativeURL); err != nil {
	// 	log.WithFields(log.Fields{
	// 		"url":        c.Flamenco.String(),
	// 		log.ErrorKey: err,
	// 	}).Error("unable to construct URL to get JWT public keys")
	// } else {
	// 	c.JWT.PublicKeysURL = jwtURL.String()
	// }
}

// GetTestConfig returns the configuration for unit tests.
// The config is loaded from `test-flamenco-manager.yaml` in the directory
// containing the caller's source.
func GetTestConfig() Conf {
	_, myFilename, _, _ := runtime.Caller(1)
	myDir := path.Dir(myFilename)

	filepath := path.Join(myDir, "test-flamenco-manager.yaml")
	conf, err := loadConf(filepath)
	if err != nil {
		log.Fatal().Err(err).Str("file", filepath).Msg("unable to load test config")
	}

	return conf
}
