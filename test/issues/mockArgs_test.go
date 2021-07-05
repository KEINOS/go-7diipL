package issues_test

import (
	"os"
	"testing"
)

// mockArgs は、テストでコマンドのオプションやフラグをモックします.
//
// 変更をリカバリーする defer 用の関数を返します.
func mockArgs(t *testing.T, argsDummy []string) func() {
	t.Helper()

	// 既存引数のバックアップとリストア
	oldOsArgs := os.Args

	// Defer 用の関数
	funcDefer := func() {
		os.Args = oldOsArgs
	}

	// 引数のモック
	os.Args = append([]string{os.Args[0]}, argsDummy...)

	return funcDefer
}
