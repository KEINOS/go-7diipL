package utils

import "os"

// IsDir returns true if directory exists in the given path and not a file.
//
// It will exit with an error if the path is malformed (can't be cleaned).
func IsDir(pathFile string) bool {
	if !PathExists(pathFile) {
		return false
	}

	fileInfo, _ := os.Stat(pathFile)

	return fileInfo.IsDir()
}
