package issues_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/app"
	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
)

// Test issue #1 fix.
func TestIssue001(t *testing.T) {
	// 標準入力のダミーデータ
	userInput := "今日はいい天気ですね"

	// stdin/パイプ からの入力のモックとリカバリ用関数取得
	funcDeferSTDIN := mockSTDIN(t, userInput)
	defer funcDeferSTDIN() // モックのリカバリ

	// オプション、フラグなどの引数のモックとリカバリ用関数取得
	funcDeferArgs := mockArgs(t, []string{
		"--info",
		"ja",
		"en",
	})
	defer funcDeferArgs() // モックのリカバリ

	appTest := app.New() // 新規アプリ作成

	if err := appTest.SetEngine("deepl", t.Name()); err != nil {
		t.Fatalf("failed to set translation engine during test.\nErr Msg: %s", err)
	}
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
