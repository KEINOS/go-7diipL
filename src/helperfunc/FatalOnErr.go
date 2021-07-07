package helperfunc

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

var (
	IsErrorDummy bool // IsErrorDummy が true の場合、FatalOnErr は t.FailNow せずに標準エラー出力します.
	failNow      func(format string, args ...interface{})
)

// FatalOnErr は err が nil ではない場合に t.FailNow でテストを終了します.
//
// 基本的に MockArgs と MockSTDIN などのヘルパー関数自身のカバレッジ目的で使用します.
// 一般的なテストでは t.Fatalf を利用してください.
func FatalOnErr(t *testing.T, err error, comment ...string) {
	t.Helper()

	if err == nil {
		return
	}

	msgAdditional := strings.TrimSpace(strings.Join(comment, "\n"))

	// msgAdditional が空ではない場合、頭に改行を追加（err 自体が改行で終わらないため可読性のため）
	if msgAdditional != "" {
		msgAdditional = "\n" + msgAdditional
	}

	// 処理内容を変更可能にするため関数を変数に代入
	failNow = t.Fatalf

	// IsErrorDummy が true にセットされていた場合は標準エラー出力にエラー内容を出力する
	if IsErrorDummy {
		failNow = func(format string, args ...interface{}) {
			fmt.Fprintf(os.Stderr, format, args...)
		}

		// FatalOnErr 自身のテスト用途以外では変更すべきではないため 1 度使われたら false に強制セット
		IsErrorDummy = false
	}

	failNow("fatal error during test.\nErrMsg: %v%v", err, msgAdditional)
}
