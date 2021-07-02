package utils

import (
	"os"
	"path/filepath"
)

// GetNameExe は、現在の実行ファイル名から拡張子を除いたファイル名を返します.
func GetNameExe() string {
	pathFileExe := os.Args[0]

	return filepath.Base(pathFileExe[:len(pathFileExe)-len(filepath.Ext(pathFileExe))])
}
