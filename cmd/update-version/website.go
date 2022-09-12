package main

import (
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"
)

const websiteFile = "web/project-website/data/flamenco.yaml"

// updateWebsite changes the version number of the latest version in the project
// website.
// Returns whether the file actually changed.
func updateWebsite() bool {
	replacer := func(line string) string {
		if !strings.HasPrefix(line, "latestVersion: ") {
			return line
		}
		return fmt.Sprintf("latestVersion: %q", cliArgs.newVersion)
	}

	fileWasChanged, err := updateLines(websiteFile, replacer)
	if err != nil {
		log.Fatal().Err(err).Msg("error updating project website")
	}
	return fileWasChanged
}
