package main

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/app"
	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
)

func TestMain_exits_status1_on_fail(t *testing.T) {
	app.ForceFailRun = true
	defer func() {
		app.ForceFailRun = false
	}()

	// モックのリカバリ
	oldOsExit := utils.OsExit
	defer func() {
		utils.OsExit = oldOsExit
	}()

	var status int // Capture exit status

	// OsExit のモック
	utils.OsExit = func(code int) {
		status = code
	}

	// Run main
	out := capturer.CaptureOutput(func() {
		main()
	})

	assert.NotEqual(t, utils.SUCCESS, status, "staus should not be zero on fail")
	assert.Empty(t, out, "forcing fail should not output anything")
}
