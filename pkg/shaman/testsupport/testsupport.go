package testsupport

import (
	"testing"

	"projects.blender.org/studio/flamenco/pkg/sysinfo"
)

func SkipTestIfUnableToSymlink(t *testing.T) {
	can, err := sysinfo.CanSymlink()
	switch {
	case err != nil:
		t.Fatalf("error checking platform symlinking capabilities: %v", err)
	case !can:
		t.Skip("symlinking not possible on current platform")
	}
}
