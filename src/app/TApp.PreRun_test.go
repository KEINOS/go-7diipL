package app_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/app"
	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/stretchr/testify/assert"
)

func TestPreRun(t *testing.T) {
	argv := new(app.TFlagOptions)

	argv.IsModeDebug = true
	defer func() {
		utils.SetModeDebug(false)
	}()

	appTest := app.New()
	appTest.Argv = argv

	// Execute pre-run
	appTest.PreRun()

	// デバッグモードのチェック
	assert.True(t, utils.IsModeDebug(),
		"if argv's IsModeDebug is true then utils.IsModeDebug should be true as well")

	// ヘルプ表示のテンプレート流し込み確認
	assert.Contains(t, appTest.Argv.UsageApp, app.NameDefault, "parsed template should contain the app name")
	assert.Contains(t, appTest.Argv.UsageApp, "app コマンド", "parsed template should contain the binary name w/out ext")
}
