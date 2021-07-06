package helperfunc

import (
	"io/ioutil"
	"os"
	"testing"
)

// mockSTDIN は、テストでユーザからの標準入力をモックするヘルパー関数です.
// この関数は、モックの変更をリカバリーする defer 用の関数を返します.
//
//   // 以下は echo "foo bar" | qiitrans と同等の状態
//   dummySTDIN := "foo bar"
//   funcDefer := helperfunc.MockSTDIN(t, dummySTDIN)
//   defer funcDefer()
//   /* DO SOMETHING WITH STDIN HERE */
func MockSTDIN(t *testing.T, inputDummy string) func() {
	t.Helper()

	/* stdin のダミー用ファイルの作成 */
	content := []byte(inputDummy)

	tmpFile, err := ioutil.TempFile("", "example")
	FatalOnErr(t, err, "failed to create temporary file for I/O")

	_, err = tmpFile.Write(content)
	FatalOnErr(t, err, "failed to write temp I/O file")

	_, err = tmpFile.Seek(0, 0)
	FatalOnErr(t, err, "failed to set the offset for the next Read or Write on temp I/O file")

	// 既存標準入力のファイルポインタのバックアップとリストア
	oldStdin := os.Stdin

	// stdin のモック
	os.Stdin = tmpFile

	// Defer 用の関数
	funcDefer := func() {
		err := os.Remove(tmpFile.Name())
		FatalOnErr(t, err, "failed to remove temp file for I/O")

		os.Stdin = oldStdin
	}

	return funcDefer
}
