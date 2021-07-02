package utils

import "os"

// PathExists はパスが存在する場合に true を返します.
func PathExists(pathFile string) bool {
	_, err := os.Stat(pathFile)

	return err == nil
}
