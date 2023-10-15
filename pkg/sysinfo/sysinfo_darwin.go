package sysinfo

// SPDX-License-Identifier: GPL-3.0-or-later

// canSymlink always returns true, as symlinking on non-Windows platforms is not
// hard.
func canSymlink() (bool, error) {
	return true, nil
}

func description() (string, error) {
	// TODO: figure out how to get more info on macOS.
	return "macOS", nil
}
