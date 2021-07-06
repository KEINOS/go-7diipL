package issues_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/app"
	"github.com/Qithub-BOT/QiiTrans/src/helperfunc"
	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
)

// Test issue #1 fix.
func TestIssue001(t *testing.T) {
	// 標準入力のダミーデータ
	userInput := "今日はいい天気ですね"

	// stdin/パイプ からの入力のモックとリカバリ用関数取得
	funcDeferSTDIN := helperfunc.MockSTDIN(t, userInput)
	defer funcDeferSTDIN() // モックのリカバリ

	// オプション、フラグなどの引数のモックとリカバリ用関数取得
	funcDeferArgs := helperfunc.MockArgs(t, []string{
		"--info",
		"ja",
		"en",
	})
	defer funcDeferArgs() // モックのリカバリ

	appTest := app.New(t.Name()) // 新規アプリ作成。キャッシュ ID はテスト名.

	defer func() {
		appTest.Engine.Cache.ClearAll() // 終了後にキャッシュ削除
	}()

	// テスト実行
	out := capturer.CaptureOutput(func() {
		status := appTest.Run()

		assert.Equal(t, 0, status)
	})

	assert.Contains(t, out, "It's a beautiful day.", "it should contain the translation")
	assert.Contains(t, out, "残り文字数",
		"if --info flag was set then it should contain the remain number of chars")
}
