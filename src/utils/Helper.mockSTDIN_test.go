package utils_test

import (
	"io/ioutil"
	"os"
	"testing"
)

// mockSTDIN は標準入力のテスト用のヘルパー関数です.
//
// モック（模擬）したい標準入力の値を渡すと、ダミーのファイル・ポインタと defer 用の関数が返ってきます.
//
// 使用例
//
//   input := "sample dinput"
//   tmpFile, funcDefer := mockSTDIN(t, input)
//
//   defer funcDefer() // clean up
//
//   oldStdin := os.Stdin
//
//   defer func() { os.Stdin = oldStdin }() // Restore original Stdin
//
//   // Mock stdin
//   os.Stdin = tmpFile
//
//   result := funcSTDINSample() // os.Stdin を使った関数
func mockSTDIN(t *testing.T, inputDummy string) (*os.File, func()) {
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

	// Defer 用の関数
	funcDefer := func() {
		os.Remove(tmpFile.Name())
	}

	return tmpFile, funcDefer
}
