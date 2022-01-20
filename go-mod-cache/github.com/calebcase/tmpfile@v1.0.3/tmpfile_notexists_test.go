// Copyright 2019 Caleb Case
// +build !windows

package tmpfile

import (
	"errors"
	"os"
)

func notExists(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return errors.New("file exists")
	} else if !os.IsNotExist(err) {
		return err
	}

	return nil
}
