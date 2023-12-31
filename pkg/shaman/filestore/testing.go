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

package filestore

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"projects.blender.org/studio/flamenco/pkg/shaman/config"
)

// CreateTestStore returns a Store that can be used for unit testing.
func CreateTestStore() *Store {
	tempDir, err := ioutil.TempDir("", "shaman-filestore-test-")
	if err != nil {
		panic(err)
	}

	conf := config.Config{
		StoragePath: tempDir,
	}
	return New(conf)
}

// CleanupTestStore deletes a store returned by CreateTestStore()
func CleanupTestStore(store *Store) {
	if err := os.RemoveAll(store.baseDir); err != nil {
		panic(err)
	}
}

// MustStoreFileForTest allows a unit test to store some file in the 'stored' storage bin.
// Any error will cause a panic.
func (s *Store) MustStoreFileForTest(checksum string, filesize int64, contents []byte) {
	file, err := s.OpenForUpload(checksum, filesize)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	written, err := file.Write(contents)
	if err != nil {
		panic(err)
	}
	if written != len(contents) {
		panic("short write")
	}

	err = s.MoveToStored(checksum, filesize, file.Name())
	if err != nil {
		panic(err)
	}
}

// LinkTestFileStore creates a copy of _test_file_store in a temporary directory.
// Panics if there are any errors.
func LinkTestFileStore(cloneTo string) {
	_, myFilename, _, _ := runtime.Caller(0)
	fileStorePath := filepath.Join(filepath.Dir(filepath.Dir(myFilename)), "_test_file_store")
	now := time.Now()

	visit := func(visitPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relpath, err := filepath.Rel(fileStorePath, visitPath)
		if err != nil {
			return err
		}

		targetPath := filepath.Join(cloneTo, relpath)
		if info.IsDir() {
			return os.MkdirAll(targetPath, 0755)
		}
		err = copyFile(visitPath, targetPath)
		if err != nil {
			return err
		}
		// Make sure we always test with fresh files by default.
		return os.Chtimes(targetPath, now, now)
	}
	if err := filepath.Walk(fileStorePath, visit); err != nil {
		panic(err)
	}
}

func copyFile(sourcePath, destPath string) error {
	// Open the source file.
	srcFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("could not open %q: %w", sourcePath, err)
	}
	defer srcFile.Close()

	// Create the destination file.
	destFile, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("could not create %q: %w", destPath, err)
	}
	defer destFile.Close()

	// Copy the contents from source to destination.
	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return fmt.Errorf("could not copy contents of %q to %q: %w", sourcePath, destPath, err)
	}

	// Flush any buffered data to ensure completion.
	err = destFile.Sync()
	if err != nil {
		return fmt.Errorf("could not sync buffer of %q to disk: %w", destPath, err)
	}

	return nil
}
