// Copyright 2019 Caleb Case

package tmpfile

import (
	"errors"
	"os"
)

// On windows, in addition to the normal path where the file has been removed
// from the directory entries, files that are pending deletion can return
// ERROR_ACCESS_DENIED.
func notExists(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return errors.New("file exists")
	} else if !(os.IsNotExist(err) || os.IsPermission(err)) {
		return err
	}

	return nil
}
