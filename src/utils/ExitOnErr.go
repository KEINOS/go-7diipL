package utils

import (
	"fmt"
	"os"
)

// ExitOnErr は err がエラーの場合のみ、エラー内容を標準エラー出力に出力して OsExit(1) を呼び出します.
func ExitOnErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		OsExit(FAILURE)
	}
}
