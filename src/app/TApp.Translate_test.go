package app_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/app"
	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
)

func TestTranslate(t *testing.T) {
	appTest := app.New(t.Name())

	err := appTest.SetEngine("deepl")
	assert.NoError(t, err)

	defer func() {
		appTest.Engine.Cache.ClearAll()
	}()

	// 翻訳する順番 日本語 -> 英語 -> 日本語
	orderLang := []string{
		"ja",
		"en",
		"ja",
	}

	// 通常テスト
	{
		// タイミングによって異なるため複数パターンをチェック
		expect := []string{
			"今日はとてもいい天気です。",
			"今日はとてもいい天気だ。",
		}

		actual, err := appTest.Translate(orderLang, "今日はとてもいい天気です。")

		assert.NoError(t, err)
		assert.Contains(t, expect, actual)
	}

	// キャッシュテスト
	{
		utils.SetModeDebug(true)

		defer func() {
			utils.SetModeDebug(false)
		}()

		out := capturer.CaptureOutput(func() {
			// タイミングによって異なるため複数パターンをチェック
			expect := []string{
				"今日はとてもいい天気です。",
				"今日はとてもいい天気だ。",
			}

			actual, err := appTest.Translate(orderLang, "今日はとてもいい天気です。")

			assert.NoError(t, err)
			assert.Contains(t, expect, actual)
		})

		assert.Contains(t, out, "EN -> JA: キャッシュ:",
			"it should output cached translation on debug mode")
	}
}
