package sysinfo

// SPDX-License-Identifier: GPL-3.0-or-later

import (
	"fmt"

	the_other_sysinfo "github.com/zcalusic/sysinfo"
)

// canSymlink always returns true, as symlinking on non-Windows platforms is not
// hard.
func canSymlink() (bool, error) {
	return true, nil
}

func description() (string, error) {
	var si the_other_sysinfo.SysInfo
	si.GetSysInfo()
	description := fmt.Sprintf("%s (%s)", si.OS.Name, si.Kernel.Release)
	return description, nil
}
