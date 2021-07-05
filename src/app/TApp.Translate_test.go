package app_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/app"
	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
)

func TestTranslate(t *testing.T) {
	appTest := app.New()

	err := appTest.SetEngine("deepl", t.Name())
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
		out, err := appTest.Translate(orderLang, "私は、賛成の反対に同意なのだ")

		assert.NoError(t, err)

		expect := "同意しないことに同意します。"
		actual := out

		assert.Equal(t, expect, actual)
	}

	// キャッシュテスト
	{
		utils.SetModeDebug(true)

		defer func() {
			utils.SetModeDebug(false)
		}()

		out := capturer.CaptureOutput(func() {
			out, err := appTest.Translate(orderLang, "私は、賛成の反対に同意なのだ")

			assert.NoError(t, err)

			expect := "同意しないことに同意します。"
			actual := out

			assert.Equal(t, expect, actual)
		})

		assert.Contains(t, out, "EN -> JA: キャッシュ: 同意しないことに同意します。",
			"it should output cached translation on debug mode")
	}
}
