package app_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/app"
	"github.com/Qithub-BOT/QiiTrans/src/helperfunc"
	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/kami-zh/go-capturer"
	"github.com/mkideal/cli"
	"github.com/stretchr/testify/assert"
)

func TestCliRun_fail_set_arg_value(t *testing.T) {
	appTest := app.New(t.Name())
	ctxDummy := new(cli.Context)

	err := appTest.CliRun(ctxDummy)

	assert.Error(t, err, "malformed context should return an error")
}

func TestCliRun_fail_prerun(t *testing.T) {
	dummySTDIN := "Hello, world!"
	dummyArgs := []string{
		"en",
		"ja",
		"es",
	}

	funcDeferSTDIN := helperfunc.MockSTDIN(t, dummySTDIN)
	defer funcDeferSTDIN()

	funcDeferArgs := helperfunc.MockArgs(t, dummyArgs)
	defer funcDeferArgs()

	appTest := app.New(t.Name())

	app.ForceFailPreRun = true
	defer func() { app.ForceFailPreRun = false }()

	out := capturer.CaptureOutput(func() {
		status := appTest.Run()

		assert.Equal(t, utils.FAILURE, status, "it should be a non-zero on pre-run fail")
	})

	assert.Contains(t, out, "PreRun was forced to fail")
}
