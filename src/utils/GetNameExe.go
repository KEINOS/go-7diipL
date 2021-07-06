package utils

import (
	"os"
	"path/filepath"
)

// GetNameExe は、コマンド名（現在の実行ファイル名からパスと拡張子を除いたファイル名）を返します.
func GetNameExe() string {
	pathFileExe := os.Args[0]

	return filepath.Base(pathFileExe[:len(pathFileExe)-len(filepath.Ext(pathFileExe))])
}
