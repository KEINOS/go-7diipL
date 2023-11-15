package app_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/app"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetUniformedInfo(t *testing.T) {
	appTest := app.New("", t.Name())

	err := appTest.SetEngine("deepl")
	require.NoError(t, err, "failed set engine during test")

	defer appTest.Engine.Cache.ClearAll()

	info, err := appTest.GetUniformedInfo()

	require.NoError(t, err)
	assert.Contains(t, info, "[INFO]: 残り文字数:")
}

func TestGetUniformedInfo_error(t *testing.T) {
	appTest := app.New("", t.Name())

	err := appTest.SetEngine("deepl")
	require.NoError(t, err, "failed set engine during test")

	defer appTest.Engine.Cache.ClearAll()

	// テスト用にアクセストークンのバックアップとリカバリ
	const nameKeyEnv = "DEEPL_API_KEY"

	// アクセストークンを一時削除
	t.Setenv(nameKeyEnv, "")

	// 実行
	info, err := appTest.GetUniformedInfo()

	// テスト
	require.Error(t, err, "missing API key should return an error")
	assert.Empty(t, info, "on error info should be empty")
}

func TestGetUniformedInfo_error_forced(t *testing.T) {
	appTest := app.New("", t.Name())

	if err := appTest.SetEngine("deepl"); err != nil {
		t.Fatalf("failed to set engine during test")
	}
	defer appTest.Engine.Cache.ClearAll()

	app.ForceErrorGetUniformedInfo = true
	defer func() { app.ForceErrorGetUniformedInfo = false }()

	info, err := appTest.GetUniformedInfo()

	require.Error(t, err, "it should return error if ForceErrorGetUniformedInfo is true")
	assert.Empty(t, info, "on error it should be empty")
}
