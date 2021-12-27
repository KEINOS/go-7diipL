package issues_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/app"
	"github.com/Qithub-BOT/QiiTrans/src/helperfunc"
	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
)

// Test issue #8 fix.
func TestIssue008(t *testing.T) {
	// 標準入力のダミーデータ
	userInput := ""

	// stdin/パイプ からの入力のモックとリカバリ用関数取得
	funcDeferSTDIN := helperfunc.MockSTDIN(t, userInput)
	defer funcDeferSTDIN() // モックのリカバリ

	// オプション、フラグなどの引数のモックとリカバリ用関数取得
	funcDeferArgs := helperfunc.MockArgs(t, []string{
		"--no-cache",
		"--debug",
		"ja",
		"en",
	})
	defer funcDeferArgs() // モックのリカバリ

	appTest := app.New("", t.Name()) // 新規アプリ作成。キャッシュ ID はテスト名.

	defer func() {
		appTest.Engine.Cache.ClearAll() // 終了後にキャッシュ削除
	}()

	// テスト実行
	out := capturer.CaptureOutput(func() {
		status := appTest.Run()

		assert.Equal(t, 0, status)
	})

	assert.NotContains(t, out, "badger", "when empty it should not call cache")
	assert.NotContains(t, out, "JA -> EN: 新規取得:", "when empty it should not request API")
}
