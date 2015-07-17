package helper

import (
	"os"
)

func IsDirExists(path string) bool {
	fi, err := os.Stat(path)

	if err != nil {
		return os.IsExist(err)
	} else {
		return fi.IsDir()
	}
	return false
}

func MkdirNotExist(path string) error {
	if !IsDirExists(path) {
		return os.MkdirAll(path, 0755)
	}
	return nil
}
