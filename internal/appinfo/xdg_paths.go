package appinfo

// SPDX-License-Identifier: GPL-3.0-or-later

import (
	"os"
	"path"
	"path/filepath"

	"github.com/adrg/xdg"
)

// customHome can be set at link time to specify the home directory for the worker.
// This can be overruled at runtime by setting the FLAMENCO_HOME enviroment variable.
// Only used in InFlamencoHome() function.
var customHome = ""

// InFlamencoHome returns the filename in the 'flamenco home' dir, and ensures
// that the directory exists.
func InFlamencoHome(filename string) (string, error) {
	flamencoHome := customHome
	if envHome, ok := os.LookupEnv("FLAMENCO_HOME"); ok {
		flamencoHome = envHome
	}
	if flamencoHome == "" {
		return xdg.DataFile(path.Join(xdgApplicationName, filename))
	}
	if err := os.MkdirAll(flamencoHome, os.ModePerm); err != nil {
		return "", err
	}
	return filepath.Join(flamencoHome, filename), nil
}
