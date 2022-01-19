package app_test

import (
	"os"
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/app"
	"github.com/Qithub-BOT/QiiTrans/src/helperfunc"
	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
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

	appTest := app.New("", t.Name())

	out := capturer.CaptureOutput(func() {
		status := appTest.Run()

		assert.Equal(t, status, utils.SUCCESS, "it should end with status zero (SUCCESS)")
	})

	assert.Contains(t, out, "¡Hola, mundo!")

	// Clean Up
	appTest.Engine.Cache.ClearAll()
}

func TestNew_verbose(t *testing.T) {
	dummySTDIN := "Hello, world!"
	dummyArgs := []string{
		"en",
		"ja",
		"es",
		"--verbose",
	}

	funcDeferSTDIN := helperfunc.MockSTDIN(t, dummySTDIN)
	defer funcDeferSTDIN()

	funcDeferArgs := helperfunc.MockArgs(t, dummyArgs)
	defer funcDeferArgs()

	appTest := app.New("", t.Name())

	out := capturer.CaptureOutput(func() {
		status := appTest.Run()

		assert.Equal(t, status, utils.SUCCESS, "it should end with status zero (SUCCESS)")
	})

	assert.Contains(t, out, "EN -> JA")
	assert.Contains(t, out, "¡Hola, mundo!")

	// Clean Up
	appTest.Engine.Cache.ClearAll()
}

func TestNew_clear_cache(t *testing.T) {
	dummySTDIN := "Hello, world!"
	dummyArgs := []string{
		"en",
		"es",
		"--clear",
	}

	funcDeferSTDIN := helperfunc.MockSTDIN(t, dummySTDIN)
	defer funcDeferSTDIN()

	funcDeferArgs := helperfunc.MockArgs(t, dummyArgs)
	defer funcDeferArgs()

	appTest := app.New("", t.Name())

	out := capturer.CaptureOutput(func() {
		status := appTest.Run()

		assert.Equal(t, status, utils.SUCCESS, "it should end with status zero (SUCCESS)")
	})

	assert.Contains(t, out, "¡Hola, mundo!")

	assert.True(t, appTest.Argv.IsNoCache, "--clear should set IsNoCache as true as well")

	// Clean Up
	appTest.Engine.Cache.ClearAll()
}

func TestNew_default_settings(t *testing.T) {
	appTest := app.New("", t.Name())

	// Check default name
	assert.Equal(t, app.NameDefault, appTest.Name, "it should be the default app name")

	// Check default version
	assert.Empty(t, appTest.Version, "version should be empty by default to auto detect")
}

func TestNew_fail_get_info(t *testing.T) {
	dummySTDIN := "Hello, world!"
	dummyArgs := []string{
		"en",
		"es",
		"--info",
	}

	funcDeferSTDIN := helperfunc.MockSTDIN(t, dummySTDIN)
	defer funcDeferSTDIN()

	funcDeferArgs := helperfunc.MockArgs(t, dummyArgs)
	defer funcDeferArgs()

	// GetSTDIN を強制的に失敗させる
	app.ForceErrorGetUniformedInfo = true
	defer func() { app.ForceErrorGetUniformedInfo = false }()

	appTest := app.New("", t.Name())

	out := capturer.CaptureOutput(func() {
		status := appTest.Run()

		assert.Equal(t, status, utils.FAILURE, "reading failure of API info should return status non-zero")
	})

	assert.Contains(t, out, "forced to return error")

	// Clean Up
	appTest.Engine.Cache.ClearAll()
}

func TestNew_fail_read_stdin(t *testing.T) {
	dummySTDIN := "Hello, world!"
	dummyArgs := []string{
		"en",
		"es",
	}

	funcDeferSTDIN := helperfunc.MockSTDIN(t, dummySTDIN)
	defer funcDeferSTDIN()

	funcDeferArgs := helperfunc.MockArgs(t, dummyArgs)
	defer funcDeferArgs()

	// GetSTDIN を強制的に失敗させる
	utils.ForceErrorGetSTDIN = true
	defer func() { utils.ForceErrorGetSTDIN = false }()

	appTest := app.New("", t.Name())

	out := capturer.CaptureOutput(func() {
		status := appTest.Run()

		assert.Equal(t, status, utils.FAILURE, "reading failure of STDIN should return status non-zero")
	})

	assert.Contains(t, out, "forced to return error")

	// Clean Up
	appTest.Engine.Cache.ClearAll()
}

func TestNew_instantiation(t *testing.T) {
	appTest1 := app.New("")
	appTest2 := app.New("")

	assert.IsType(t, appTest1, appTest2, "it should create the same type")
	assert.NotSame(t, appTest1, appTest2, "pointers should not reference the same object")
}

func TestNew_no_apikey_set(t *testing.T) {
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

	// テスト用にアクセストークンのバックアップとリカバリ
	nameKeyEnv := "DEEPL_API_KEY"
	oldKey := os.Getenv(nameKeyEnv)

	defer func() {
		err := os.Setenv(nameKeyEnv, oldKey)
		if err != nil {
			t.Fatalf("failed to recover env key during test")
		}
	}()

	appTest := app.New("", t.Name())

	// アクセストークンの環境変数を一時的に変更
	if err := os.Setenv(nameKeyEnv, ""); err != nil {
		t.Fatalf("failed to change env key during test")
	}

	out := capturer.CaptureOutput(func() {
		t.Logf("Current Env: %v", os.Getenv(nameKeyEnv))
		status := appTest.Run()

		assert.Greater(t, status, 0, "missing api key should be a non-zero")
	})

	assert.Contains(t, out, "API key for DeepL not set")

	// Clean Up
	appTest.Engine.Cache.ClearAll()
}

func TestNew_show_info(t *testing.T) {
	dummySTDIN := "Hello, world!"
	dummyArgs := []string{
		"en",
		"ja",
		"es",
		"--info",
	}

	funcDeferSTDIN := helperfunc.MockSTDIN(t, dummySTDIN)
	defer funcDeferSTDIN()

	funcDeferArgs := helperfunc.MockArgs(t, dummyArgs)
	defer funcDeferArgs()

	appTest := app.New("", t.Name())

	out := capturer.CaptureOutput(func() {
		status := appTest.Run()

		assert.Equal(t, status, utils.SUCCESS, "it should end with status zero (SUCCESS)")
	})

	assert.Contains(t, out, "[INFO]: 残り文字数:")

	// Clean Up
	appTest.Engine.Cache.ClearAll()
}
