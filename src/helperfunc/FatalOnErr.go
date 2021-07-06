package helperfunc

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

var (
	IsErrorDummy bool // FatalOnErr 自身のテストのため、動作をモックするためのフラグです.
	failNow      func(format string, args ...interface{})
)

// FatalOnErr は err が nil ではない場合に t.FailNow でテストを終了します.
//
// 基本的に MockArgs と MockSTDIN などのヘルパー関数自身のカバレッジ目的で使用します.
// 一般的なテストでは t.Fatalf を利用してください.
func FatalOnErr(t *testing.T, err error, comment ...string) {
	t.Helper()

	failNow = t.Fatalf

	msgAdditional := strings.TrimSpace(strings.Join(comment, "\n"))

	if msgAdditional != "" {
		msgAdditional = "\n" + msgAdditional
	}

	// IsErrorDummy が true にセットされていた場合は標準エラー出力にエラー内容を出力します.
	if IsErrorDummy {
		failNow = func(format string, args ...interface{}) {
			fmt.Fprintf(os.Stderr, format, args...)
		}

		// FatalOnErr 自身のテスト用途以外では変更すべきではないため 1 度使われたら false に強制セット
		IsErrorDummy = false
	}

	if err != nil {
		failNow("fatal error during test.\nErrMsg: %v%v", err, msgAdditional)
	}
}
