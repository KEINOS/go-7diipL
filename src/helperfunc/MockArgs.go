package helperfunc

import (
	"os"
	"testing"
)

// mockArgs は、テストでユーザのコマンド引数（オプションやフラグ含む入力）をモックするヘルパー関数です.
// この関数は、モックの変更をリカバリーする defer 用の関数を返します.
//
//   dummyArgs := []string{"--debug", "ja", "en"}
//   funcDefer := helperfunc.MockArgd(t, dummyArgs)
//   defer funcDefer()
//   /* DO SOMETHING WITH ARGS HERE */
func MockArgs(t *testing.T, argsDummy []string) func() {
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
