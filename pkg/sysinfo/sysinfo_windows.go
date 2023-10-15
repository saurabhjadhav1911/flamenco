package sysinfo

// SPDX-License-Identifier: GPL-3.0-or-later

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"golang.org/x/sys/windows/registry"
)

// editionCanSymlink maps the EditionID key in the registry to a boolean indicating whether
// symlinks are available.
var editionCanSymlink = map[string]bool{
	"Core":         false, // Windows Home misses the tool to allow symlinking to a user.
	"Professional": true,  // Still requires explicit right to be assigned to the user.
}

// canSymlink tries to determine whether the running system can use symbolic
// links, based on information from the Windows registry.
func canSymlink() (bool, error) {
	editionID, err := windowsEditionID()
	if err != nil {
		return false, fmt.Errorf("determining edition of Windows: %w", err)
	}

	canSymlink, found := editionCanSymlink[editionID]
	if !found {
		log.Warn().Str("editionID", editionID).Msg("unknown Windows edition, assuming it can use symlinks")
		return true, nil
	}

	return canSymlink, nil
}

func description() (string, error) {
	productName, err := registryReadString(
		`SOFTWARE\Microsoft\Windows NT\CurrentVersion`,
		"ProductName",
	)
	if err != nil {
		return "", err
	}

	editionID, err := windowsEditionID()
	if err != nil {
		return "", err
	}

	description := fmt.Sprintf("%s (%s)", productName, editionID)
	return description, nil
}

func windowsEditionID() (string, error) {
	// Values seen so far:
	// - "Professional"
	// - "Core"
	return registryReadString(
		`SOFTWARE\Microsoft\Windows NT\CurrentVersion`,
		"EditionID",
	)
}

func registryReadString(keyPath, valueName string) (string, error) {
	regkey, err := registry.OpenKey(
		registry.LOCAL_MACHINE,
		keyPath,
		registry.QUERY_VALUE)
	if err != nil {
		return "", fmt.Errorf("opening registry key %q: %w", keyPath, err)
	}
	defer regkey.Close()

	value, _, err := regkey.GetStringValue(valueName)
	if err != nil {
		return "", fmt.Errorf("reading registry key %q: %w", valueName, err)
	}

	return value, nil
}
