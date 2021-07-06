package app_test

import (
	"os"
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/app"
	"github.com/stretchr/testify/assert"
)

func TestGetUniformedInfo(t *testing.T) {
	appTest := app.New(t.Name())

	if err := appTest.SetEngine("deepl"); err != nil {
		t.Fatalf("failed to set engine during test")
	}
	defer appTest.Engine.Cache.ClearAll()

	info, err := appTest.GetUniformedInfo()

	assert.NoError(t, err)
	assert.Contains(t, info, "[INFO]: 残り文字数:")
}

func TestGetUniformedInfo_error(t *testing.T) {
	appTest := app.New(t.Name())

	if err := appTest.SetEngine("deepl"); err != nil {
		t.Fatalf("failed set engine during test")
	}
	defer appTest.Engine.Cache.ClearAll()

	// テスト用にアクセストークンのバックアップとリカバリ
	nameKeyEnv := "DEEPL_API_KEY"
	oldKey := os.Getenv(nameKeyEnv)

	defer func() {
		err := os.Setenv(nameKeyEnv, oldKey)
		if err != nil {
			t.Fatalf("failed to recover env key during test")
		}
	}()

	// アクセストークンを一時削除
	if err := os.Setenv(nameKeyEnv, ""); err != nil {
		t.Fatalf("failed to recover env key during test")
	}

	// 実行
	info, err := appTest.GetUniformedInfo()

	// テスト
	assert.Error(t, err, "missing API key should return an error")
	assert.Empty(t, info, "on error info should be empty")
}

func TestGetUniformedInfo_error_forced(t *testing.T) {
	appTest := app.New(t.Name())

	if err := appTest.SetEngine("deepl"); err != nil {
		t.Fatalf("failed to set engine during test")
	}
	defer appTest.Engine.Cache.ClearAll()

	app.ForceErrorGetUniformedInfo = true
	defer func() { app.ForceErrorGetUniformedInfo = false }()

	info, err := appTest.GetUniformedInfo()

	assert.Error(t, err, "it should return error if ForceErrorGetUniformedInfo is true")
	assert.Empty(t, info, "on error it should be empty")
}
