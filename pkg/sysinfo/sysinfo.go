package sysinfo

// SPDX-License-Identifier: GPL-3.0-or-later

// CanSymlink tries to determine whether the running system can use symbolic
// links.
func CanSymlink() (bool, error) {
	return canSymlink()
}

// Description returns a string that describes the current platform in more
// detail than runtime.GOOS does.
func Description() (string, error) {
	return description()
}
