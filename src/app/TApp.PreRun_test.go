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
	argv.NameEngine = "deepl"

	defer func() { utils.SetModeDebug(false) }()

	appTest := app.New(t.Name())
	appTest.Argv = argv

	// Execute pre-run
	err := appTest.PreRun()
	assert.NoError(t, err)

	// デバッグモードのチェック
	assert.True(t, utils.IsModeDebug(),
		"if argv's IsModeDebug is true then utils.IsModeDebug should be true as well")

	// ヘルプ表示のテンプレート流し込み確認
	assert.Contains(t, appTest.Argv.UsageApp, app.NameDefault, "parsed template should contain the app name")
	assert.Contains(t, appTest.Argv.UsageApp, "app コマンド", "parsed template should contain the binary name w/out ext")
}

func TestPreRun_bad_engine_name(t *testing.T) {
	argv := new(app.TFlagOptions)

	argv.NameEngine = "unknown"

	appTest := app.New(t.Name())
	appTest.Argv = argv

	// Execute pre-run
	err := appTest.PreRun()

	assert.Error(t, err)
}

func TestPreRun_force_fail(t *testing.T) {
	appTest := app.New(t.Name())
	argv := new(app.TFlagOptions)

	argv.NameEngine = "deepl"
	appTest.Argv = argv

	app.ForceFailPreRun = true
	defer func() { app.ForceFailPreRun = false }()

	// Execute pre-run
	err := appTest.PreRun()
	assert.Error(t, err)
}
