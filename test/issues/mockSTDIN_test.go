package issues_test

import (
	"io/ioutil"
	"os"
	"testing"
)

// mockSTDIN は標準入力のテスト用のヘルパー関数です.
//
// モック（模擬）したい標準入力の値を渡すと、標準入力 defer 用の関数が返ってきます.
//
// 使用例
//
//   input := "sample input"
//   uncDefer := mockSTDIN(t, input)
//
//   defer funcDefer() // recover
//
//   result := funcSTDINSample() // os.Stdin を使った関数
func mockSTDIN(t *testing.T, inputDummy string) func() {
	t.Helper()

	/* stdin のダミー用ファイルの作成 */
	content := []byte(inputDummy)

	tmpFile, err := ioutil.TempFile("", "example")
	if err != nil {
		t.Fatal("failed to create temp file during test")
	}

	if _, err := tmpFile.Write(content); err != nil {
		t.Fatal("failed to write temp file during test")
	}

	if _, err := tmpFile.Seek(0, 0); err != nil {
		t.Fatal("failed to set the offset for the next Read during test")
	}

	// 既存標準入力のファイルポインタのバックアップとリストア
	oldStdin := os.Stdin

	// stdin のモック
	os.Stdin = tmpFile

	// Defer 用の関数
	funcDefer := func() {
		err := os.Remove(tmpFile.Name())
		if err != nil {
			t.Fatalf("failed to remove temp file for io during test")
		}

		os.Stdin = oldStdin
	}

	return funcDefer
}
