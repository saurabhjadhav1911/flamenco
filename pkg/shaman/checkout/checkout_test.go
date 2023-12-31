package checkout

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"projects.blender.org/studio/flamenco/pkg/api"
	"projects.blender.org/studio/flamenco/pkg/shaman/filestore"
	"projects.blender.org/studio/flamenco/pkg/shaman/testsupport"
)

func TestCheckout(t *testing.T) {
	testsupport.SkipTestIfUnableToSymlink(t)

	manager, cleanup := createTestManager()
	defer cleanup()
	ctx := context.Background()

	filestore.LinkTestFileStore(manager.fileStore.BasePath())

	checkout := api.ShamanCheckout{
		CheckoutPath: "á hausinn á þér",
		Files: []api.ShamanFileSpec{
			{Sha: "590c148428d5c35fab3ebad2f3365bb469ab9c531b60831f3e826c472027a0b9", Size: 3367, Path: "subdir/replacer.py"},
			{Sha: "80b749c27b2fef7255e7e7b3c2029b03b31299c75ff1f1c72732081c70a713a3", Size: 7488, Path: "feed.py"},
			{Sha: "914853599dd2c351ab7b82b219aae6e527e51518a667f0ff32244b0c94c75688", Size: 486, Path: "httpstuff.py"},
			{Sha: "d6fc7289b5196cc96748ea72f882a22c39b8833b457fe854ef4c03a01f5db0d3", Size: 7217, Path: "много ликова.py"},
		},
	}

	actualCheckoutPath, err := manager.Checkout(ctx, checkout)
	if err != nil {
		t.Fatalf("fatal error: %v", err)
	}

	// Check the symlinks of the checkout
	coPath := filepath.Join(manager.checkoutBasePath, actualCheckoutPath)
	assert.FileExists(t, filepath.Join(coPath, "subdir", "replacer.py"))
	assert.FileExists(t, filepath.Join(coPath, "feed.py"))
	assert.FileExists(t, filepath.Join(coPath, "httpstuff.py"))
	assert.FileExists(t, filepath.Join(coPath, "много ликова.py"))

	storePath := manager.fileStore.StoragePath()
	assertLinksTo(t, filepath.Join(coPath, "subdir", "replacer.py"),
		filepath.Join(storePath, "59", "0c148428d5c35fab3ebad2f3365bb469ab9c531b60831f3e826c472027a0b9", "3367.blob"))
	assertLinksTo(t, filepath.Join(coPath, "feed.py"),
		filepath.Join(storePath, "80", "b749c27b2fef7255e7e7b3c2029b03b31299c75ff1f1c72732081c70a713a3", "7488.blob"))
	assertLinksTo(t, filepath.Join(coPath, "httpstuff.py"),
		filepath.Join(storePath, "91", "4853599dd2c351ab7b82b219aae6e527e51518a667f0ff32244b0c94c75688", "486.blob"))
	assertLinksTo(t, filepath.Join(coPath, "много ликова.py"),
		filepath.Join(storePath, "d6", "fc7289b5196cc96748ea72f882a22c39b8833b457fe854ef4c03a01f5db0d3", "7217.blob"))
}

func assertLinksTo(t *testing.T, linkPath, expectedTarget string) {
	actualTarget, err := os.Readlink(linkPath)
	assert.NoError(t, err)
	assert.Equal(t, expectedTarget, actualTarget)
}

func Test_isValidCheckoutPath(t *testing.T) {
	tests := []struct {
		name         string
		checkoutPath string
		want         bool
	}{
		// Valid cases.
		{"simple", "a", true},
		{"uuid", "5e5be786-e6d7-480c-90e6-437f9ef5bf5d", true},
		{"with-spaces", "5e5be786 e6d7 480c 90e6 437f9ef5bf5d", true},
		{"project-scene-job-discriminator", "Sprite-Fright/scenename/jobname/2022-03-25-11-30-feb3", true},
		{"unicode", "ránið/lélegt vélmenni", true},

		// Invalid cases.
		{"empty", "", false},
		{"backslashes", "with\\backslash", false},
		{"windows-drive-letter", "c:/blah", false},
		{"question-mark", "blah?", false},
		{"star", "blah*hi", false},
		{"semicolon", "blah;hi", false},
		{"colon", "blah:hi", false},
		{"absolute-path", "/blah", false},
		{"directory-up", "path/../../../../etc/passwd", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidCheckoutPath(tt.checkoutPath); got != tt.want {
				t.Errorf("isValidCheckoutPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
