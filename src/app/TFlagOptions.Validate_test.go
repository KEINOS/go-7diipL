package app_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/app"
	"github.com/Qithub-BOT/QiiTrans/src/helperfunc"
	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
)

func TestValidate_flag_is_version(t *testing.T) {
	funcDefer := helperfunc.MockArgs(t, []string{
		"--version",
	})
	defer funcDefer()

	appTest := app.New(t.Name())

	out := capturer.CaptureOutput(func() {
		status := appTest.Run()

		assert.Equal(t, utils.SUCCESS, status, "when no failure it should return zero(SUCCESS)")
	})

	assert.Contains(t, out, "dev version")

	// Clean Up
	defer appTest.Engine.Cache.ClearAll()
}

func TestValidate_all_args_missing(t *testing.T) {
	dummySTDIN := "Hello, world"
	missingAllArgs := []string{}

	funcDeferSTDIN := helperfunc.MockSTDIN(t, dummySTDIN)
	defer funcDeferSTDIN()

	funcDeferArgs := helperfunc.MockArgs(t, missingAllArgs)
	defer funcDeferArgs()

	appTest := app.New(t.Name())

	out := capturer.CaptureOutput(func() {
		status := appTest.Run()

		assert.Greater(t, status, 0, "missing arg should be non-zero status")
	})

	assert.Contains(t, out, "引数が足りません。最低でも翻訳元と翻訳先の言語を指定してください")
}

func TestValidate_missing_2nd_arg(t *testing.T) {
	dummySTDIN := "Hello, world"
	missingArg := []string{"ja"} // 2 args are required

	funcDeferSTDIN := helperfunc.MockSTDIN(t, dummySTDIN)
	defer funcDeferSTDIN()

	funcDeferArgs := helperfunc.MockArgs(t, missingArg)
	defer funcDeferArgs()

	appTest := app.New(t.Name())

	out := capturer.CaptureOutput(func() {
		status := appTest.Run()

		assert.Greater(t, status, 0, "missing arg should be non-zero status")
	})

	assert.Contains(t, out, "引数が足りません。翻訳先の言語を指定してください")
}

func TestValidate_bad_lang_type(t *testing.T) {
	dummySTDIN := "Hello, world"
	badLang := "unknownlangtype"

	funcDeferSTDIN := helperfunc.MockSTDIN(t, dummySTDIN)
	defer funcDeferSTDIN()

	funcDeferArgs := helperfunc.MockArgs(t, []string{
		"ja",
		badLang,
	})
	defer funcDeferArgs()

	appTest := app.New(t.Name())

	out := capturer.CaptureOutput(func() {
		status := appTest.Run()

		assert.Greater(t, status, 0, "missing arg should be non-zero status")
	})

	assert.Contains(t, out, "引数の言語名が未定義のものです: "+badLang)
}
