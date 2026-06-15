package app_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/app"
	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	// inputJapanese is the Japanese sentence used as translation input in tests.
	inputJapanese = "今日はとてもいい天気です。"
)

//nolint:paralleltest // due to the monkey patching of global variable(s)
func TestTranslate(t *testing.T) {
	requireDeepLAPIKey(t)

	appTest := app.New("", t.Name())

	err := appTest.SetEngine("deepl")
	require.NoError(t, err)

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
		// タイミングによって結果が異なるため複数パターンをチェック
		expect := []string{
			inputJapanese,
			"今日はとてもいい天気だ。",
			"今日はとてもいい天気ですね。",
			"今日はとてもいい日だ。",
		}

		listTranslated, err := appTest.Translate(orderLang, inputJapanese)
		require.NoError(t, err)
		require.NotEmpty(t, listTranslated)

		actual := listTranslated[len(listTranslated)-1].Translated
		assert.Contains(t, expect, actual)
	}

	// キャッシュテスト
	{
		utils.SetModeDebug(true)

		defer func() {
			utils.SetModeDebug(false)
		}()

		out := capturer.CaptureOutput(func() {
			// タイミングによって結果が異なるため複数パターンをチェック
			expect := []string{
				inputJapanese,
				"今日はとてもいい天気だ。",
				"今日はとてもいい天気ですね。",
				"今日はとてもいい日だ。",
			}

			listTranslated, err := appTest.Translate(orderLang, inputJapanese)
			require.NoError(t, err)
			require.NotEmpty(t, listTranslated)

			actual := listTranslated[len(listTranslated)-1].Translated
			assert.Contains(t, expect, actual)
		})

		assert.Contains(t, out, "EN -> JA: キャッシュ:",
			"it should output cached translation on debug mode")
	}
}

func TestTranslate_force_not_translate(t *testing.T) {
	t.Parallel()

	appTest := app.New("", t.Name())
	orderLang := []string{"ja", "en", "ja"}

	appTest.Force["NoTrans"] = true

	expect := inputJapanese

	listTranslated, err := appTest.Translate(orderLang, expect)
	require.NoError(t, err)
	require.NotEmpty(t, listTranslated)

	actual := listTranslated[len(listTranslated)-1].Translated
	assert.Equal(t, expect, actual)
}

func TestTranslate_force_translate_error(t *testing.T) {
	t.Parallel()

	appTest := app.New("", t.Name())
	orderLang := []string{"ja", "en", "ja"}

	appTest.Force["TransError"] = true

	input := inputJapanese
	listTranslated, err := appTest.Translate(orderLang, input)

	require.Error(t, err)
	require.Nil(t, listTranslated)
}
