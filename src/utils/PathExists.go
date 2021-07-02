package utils

import "os"

// PathExists returns true if the given path is a file and exists.
func PathExists(pathFile string) bool {
	_, err := os.Stat(pathFile)

	return err == nil
}
