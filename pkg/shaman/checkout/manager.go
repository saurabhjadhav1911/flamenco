/* (c) 2019, Blender Foundation - Sybren A. Stüvel
 *
 * Permission is hereby granted, free of charge, to any person obtaining
 * a copy of this software and associated documentation files (the
 * "Software"), to deal in the Software without restriction, including
 * without limitation the rights to use, copy, modify, merge, publish,
 * distribute, sublicense, and/or sell copies of the Software, and to
 * permit persons to whom the Software is furnished to do so, subject to
 * the following conditions:
 *
 * The above copyright notice and this permission notice shall be
 * included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
 * EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
 * MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
 * IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
 * CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
 * TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
 * SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package checkout

import (
	"errors"
	"fmt"
	"io/fs"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/rs/zerolog/log"

	"projects.blender.org/studio/flamenco/pkg/shaman/config"
	"projects.blender.org/studio/flamenco/pkg/shaman/filestore"
	"projects.blender.org/studio/flamenco/pkg/shaman/touch"
)

var (
	// ErrDoesNotExist is returned by EraseCheckout().
	ErrDoesNotExist = errors.New("checkout does not exist")
)

// Manager creates checkouts and provides info about missing files.
type Manager struct {
	checkoutBasePath string
	fileStore        *filestore.Store

	wg *sync.WaitGroup

	checkoutUniquenessMutex *sync.Mutex
}

// ResolvedCheckoutInfo contains the result of validating the Checkout ID and parsing it into a final path.
type ResolvedCheckoutInfo struct {
	// The absolute path on our filesystem.
	absolutePath string
	// The path relative to the Manager.checkoutBasePath. This is what was
	// received from the client, updated to be unique.
	RelativePath string
}

type ErrInvalidCheckoutPath struct {
	CheckoutPath string
}

func (err ErrInvalidCheckoutPath) Error() string {
	return fmt.Sprintf("invalid checkout path %q", err.CheckoutPath)
}

// Errors returned by the Checkout Manager.
var (
	ErrCheckoutAlreadyExists = errors.New("A checkout with this ID already exists")
)

// NewManager creates and returns a new Checkout Manager.
func NewManager(conf config.Config, fileStore *filestore.Store) *Manager {
	checkoutDir := conf.CheckoutPath()
	logger := log.With().Str("checkoutDir", checkoutDir).Logger()
	logger.Info().Msg("opening checkout directory")

	err := os.MkdirAll(checkoutDir, 0777)
	if err != nil {
		logger.Error().Err(err).Msg("unable to create checkout directory")
	}

	return &Manager{checkoutDir, fileStore, new(sync.WaitGroup), new(sync.Mutex)}
}

// Close waits for still-running touch() calls to finish, then returns.
func (m *Manager) Close() {
	log.Info().Msg("shutting down Checkout manager")
	m.wg.Wait()
}

func (m *Manager) pathForCheckout(requestedCheckoutPath string) (ResolvedCheckoutInfo, error) {
	if !isValidCheckoutPath(requestedCheckoutPath) {
		return ResolvedCheckoutInfo{}, ErrInvalidCheckoutPath{requestedCheckoutPath}
	}

	return ResolvedCheckoutInfo{
		absolutePath: filepath.Join(m.checkoutBasePath, requestedCheckoutPath),
		RelativePath: requestedCheckoutPath,
	}, nil
}

// PrepareCheckout creates the root directory for a specific checkout.
// Returns the path relative to the checkout root directory.
func (m *Manager) PrepareCheckout(requestedCheckoutPath string) (ResolvedCheckoutInfo, error) {
	// This function checks the filesystem and tries to ensure uniqueness, so it's
	// important that it doesn't run simultaneously in parallel threads.
	m.checkoutUniquenessMutex.Lock()
	defer m.checkoutUniquenessMutex.Unlock()

	var lastErr error
	attemptCheckoutPath := requestedCheckoutPath

	// Just try 10 different random suffixes. If that still doesn't work, fail.
	for try := 0; try < 10; try++ {
		checkoutPaths, err := m.pathForCheckout(attemptCheckoutPath)
		if err != nil {
			return ResolvedCheckoutInfo{}, err
		}

		logger := log.With().
			Str("absolutePath", checkoutPaths.absolutePath).
			Str("checkoutPath", requestedCheckoutPath).
			Logger()

		if stat, err := os.Stat(checkoutPaths.absolutePath); !errors.Is(err, fs.ErrNotExist) {
			if err == nil {
				// No error stat'ing this path, indicating it's an existing checkout.
				lastErr = ErrCheckoutAlreadyExists
				if stat.IsDir() {
					logger.Debug().Msg("shaman: checkout path exists, going to add a random suffix")
				} else {
					logger.Warn().Msg("shaman: checkout path exists but is not a directory")
				}

				// Retry with (another) random suffix.
				attemptCheckoutPath = fmt.Sprintf("%s-%s", requestedCheckoutPath, randomisedToken())
				continue
			}
			// If it's any other error, it's really a problem on our side. Don't retry.
			logger.Error().Err(err).Msg("shaman: unable to stat checkout directory")
			return ResolvedCheckoutInfo{}, err
		}

		if err := os.MkdirAll(checkoutPaths.absolutePath, 0777); err != nil {
			lastErr = err
			logger.Warn().Err(err).Msg("shaman: unable to create checkout directory")
			continue
		}

		log.Debug().
			Str("requestedPath", requestedCheckoutPath).
			Str("actualPath", checkoutPaths.RelativePath).
			Msg("shaman: created checkout directory")
		return checkoutPaths, nil
	}

	return ResolvedCheckoutInfo{}, lastErr
}

// EraseCheckout removes the checkout directory structure identified by the ID.
// Returns ErrDoesNotExist if the checkout with this ID does not exist.
func (m *Manager) EraseCheckout(checkoutID string) error {
	checkoutPaths, err := m.pathForCheckout(checkoutID)
	if err != nil {
		return err
	}
	_, err = os.Stat(checkoutPaths.absolutePath)
	switch {
	case err == nil:
		break
	case errors.Is(err, os.ErrNotExist):
		return ErrDoesNotExist
	default:
		return err
	}

	logger := log.With().
		Str("checkoutPath", checkoutPaths.absolutePath).
		Str("checkoutID", checkoutID).
		Logger()
	if err := os.RemoveAll(checkoutPaths.absolutePath); err != nil {
		logger.Error().Err(err).Msg("shaman: unable to remove checkout directory")
		return err
	}

	// Try to remove the parent path as well, to not keep the dangling two-letter dirs.
	// Failure is fine, though, because there is no guarantee it's empty anyway.
	os.Remove(filepath.Dir(checkoutPaths.absolutePath))
	logger.Info().Msg("shaman: removed checkout directory")
	return nil
}

// SymlinkToCheckout creates a symlink at symlinkPath to blobPath.
// It does *not* do any validation of the validity of the paths!
func (m *Manager) SymlinkToCheckout(blobPath, checkoutPath, symlinkRelativePath string) error {
	symlinkPath := filepath.Join(checkoutPath, symlinkRelativePath)
	logger := log.With().
		Str("blobPath", blobPath).
		Str("symlinkPath", symlinkPath).
		Logger()

	blobPath, err := filepath.Abs(blobPath)
	if err != nil {
		logger.Error().Err(err).Msg("shaman: unable to make blobPath absolute")
		return err
	}

	logger.Debug().Msg("shaman: creating symlink")

	// This is expected to fail sometimes, because we don't create parent directories yet.
	// We only create those when we get a failure from symlinking.
	err = os.Symlink(blobPath, symlinkPath)
	switch {
	case err == nil:
		return nil
	case errors.Is(err, fs.ErrExist):
		// The symlink already exists, which is weird. Investigate so we can log a
		// more detailed warning.
		linkTarget, readErr := os.Readlink(symlinkPath)
		if readErr != nil {
			logger.Error().
				AnErr("symlinkError", err).
				AnErr("symlinkReadError", readErr).
				Msg("shaman: unable to create symlink as it already exists, but also it cannot be read")
			return err
		}
		if linkTarget != blobPath {
			logger.Error().
				AnErr("symlinkError", err).
				Str("alreadyLinkedFrom", linkTarget).
				Msg("shaman: unable to create simlink, as it already exists and links a different blob")
			return err
		}
		// The right file is linked, so let's warn about the situation and otherwise ignore it.
		logger.Warn().
			AnErr("symlinkError", err).
			Msg("shaman: symlink unexpectedly already exists, but it is linking the right path so let's just use it")
	case errors.Is(err, fs.ErrNotExist):
		// The directory doesn't exist yet.
		logger.Debug().Msg("shaman: creating parent directory")
		dir := filepath.Dir(symlinkPath)
		if err := os.MkdirAll(dir, 0777); err != nil {
			logger.Error().Err(err).Msg("shaman: unable to create parent directory")
			return err
		}
		if err := os.Symlink(blobPath, symlinkPath); err != nil {
			logger.Error().Err(err).Msg("shaman: unable to create symlink, after creating parent directory")
			return err
		}
	default:
		logger.Error().Err(err).Msg("shaman: unable to create symlink")
		return err
	}

	// Change the modification time of the blob to mark it as 'referenced' just now.
	m.wg.Add(1)
	go func() {
		if err := touchFile(blobPath); err != nil {
			logger.Warn().Err(err).Msg("shaman: unable to touch blob path")
		}
		m.wg.Done()
	}()

	return nil
}

// touchFile changes the modification time of the blob to mark it as 'referenced' just now.
func touchFile(blobPath string) error {
	if blobPath == "" {
		return os.ErrInvalid
	}
	logger := log.With().Str("file", blobPath).Logger()
	logger.Debug().Msg("shaman: touching file")

	now := time.Now()
	err := touch.Touch(blobPath)
	if err != nil {
		return err
	}

	duration := time.Since(now)
	if duration > 1*time.Second {
		logger.Warn().Str("duration", duration.String()).Msg("done touching but took a long time")
	}

	logger.Debug().Msg("done touching")
	return nil
}

// randomisedToken generates a random 4-character string.
// It is intended to add to a checkout path, to create some randomness and thus
// a higher chance of the path not yet existing.
func randomisedToken() string {
	var runes = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

	n := 4
	s := make([]rune, n)
	for i := range s {
		s[i] = runes[rand.Intn(len(runes))]
	}
	return string(s)
}
