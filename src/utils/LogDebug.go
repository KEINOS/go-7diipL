package utils

import (
	"fmt"
	"os"
	"strings"
)

// LogDebug はデバッグ・モードが true の場合のみ標準エラー出力に出力します.
//
// デバッグ・モードをセットするには SetModeDebug() を使ってください.
func LogDebug(log string, logs ...interface{}) {
	// isModeDebug は SetModeDebug.go で定義してます
	if !isModeDebug {
		return
	}

	format := "[LOG]: " + strings.TrimRight(log, "\n") + "\n"

	fmt.Fprintf(os.Stderr, format, logs...)
}
