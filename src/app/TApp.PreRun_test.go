package app_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/app"
	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPreRun(t *testing.T) {
	const nameEngine = "deepl"

	argv := new(app.TFlagOptions)

	argv.IsModeDebug = true
	argv.NameEngine = nameEngine

	defer func() { utils.SetModeDebug(false) }()

	appTest := app.New("", t.Name())
	appTest.Argv = argv

	// Execute pre-run
	err := appTest.PreRun()
	require.NoError(t, err, "PreRun should not return an error")

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

	appTest := app.New("", t.Name())
	appTest.Argv = argv

	// Execute pre-run
	err := appTest.PreRun()

	require.Error(t, err, "unknown engine name should return an error")
}

func TestPreRun_force_fail(t *testing.T) {
	const nameEngine = "deepl"

	appTest := app.New("", t.Name())
	argv := new(app.TFlagOptions)

	argv.NameEngine = nameEngine
	appTest.Argv = argv

	appTest.Force["FailPreRun"] = true

	// Execute pre-run
	err := appTest.PreRun()

	require.Error(t, err, "forced error should return an error")
}

func TestPreRun_force_fail_not_piped(t *testing.T) {
	const nameEngine = "deepl"

	appTest := app.New("", t.Name())
	argv := new(app.TFlagOptions)

	argv.NameEngine = nameEngine
	appTest.Argv = argv

	appTest.Force["IsNotPiped"] = true

	// Execute pre-run
	err := appTest.PreRun()

	require.NoError(t, err, "forced error should return an error")
	assert.False(t, appTest.Argv.IsPiped)
}
