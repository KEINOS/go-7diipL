package app_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/app"
	"github.com/Qithub-BOT/QiiTrans/src/helperfunc"
	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
)

func TestAutoHelp(t *testing.T) {
	dummySTDIN := ""
	onlyHelpOption := []string{"--help"}

	funcDeferSTDIN := helperfunc.MockSTDIN(t, dummySTDIN)
	defer funcDeferSTDIN()

	funcDeferArgs := helperfunc.MockArgs(t, onlyHelpOption)
	defer funcDeferArgs()

	appTest := app.New("", t.Name())

	out := capturer.CaptureOutput(func() {
		status := appTest.Run()

		assert.Equal(t, utils.SUCCESS, status, "--help should end with status zero")
	})

	assert.Contains(t, out, appTest.Name, "help msg should contain the app name")
	assert.Contains(t, out, "app コマンドは文書作成支援ツールです。", "help msg should contain the app description")
}
