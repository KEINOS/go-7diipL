package utils

import (
	"os"
)

// IsFile returns true if file exists in the given path.
func IsFile(pathFile string) bool {
	if !PathExists(pathFile) {
		return false
	}

	fileInfo, _ := os.Stat(pathFile)

	return !fileInfo.IsDir()
}
