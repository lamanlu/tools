package common

import (
	"io/fs"
	"os"
)

// Filesystem helpers.
func touchPath(path string, mode uint32) error {
	_, err := os.Stat(path)
	if !os.IsNotExist(err) {
		return err
	}
	err = os.MkdirAll(path, fs.FileMode(mode))
	return err
}

func fileExist(file string) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}
