package utils

import (
	"fmt"
	"os"
)

// EchoSTDERR は fmt.Fprintf のラッパー関数で、標準エラー出力に出力します.
//
//nolint:goprintffuncname // EchoSTDERRf rather than EchoSTDERR is ideal, but keep it short
func EchoSTDERR(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
}
