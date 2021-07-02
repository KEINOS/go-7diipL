package utils

import "os"

// IsDir は pathFile のパスが存在し、ディレクトリの場合に true を返します.
//
// パスの構文解析に失敗した場合はエラーで終了（OsExit）します.
func IsDir(pathFile string) bool {
	if !PathExists(pathFile) {
		return false
	}

	fileInfo, err := os.Stat(pathFile)
	ExitOnErr(err)

	return fileInfo.IsDir()
}
