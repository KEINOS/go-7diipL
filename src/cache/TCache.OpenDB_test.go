package cache_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/cache"
	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
)

func TestOpenDB(t *testing.T) {
	tmpCache := cache.New(t.Name())

	utils.SetModeDebug(true) // デバッグ・モード有効

	defer func() {
		tmpCache.ClearAll()       // 終了時にキャッシュを削除
		utils.SetModeDebug(false) // 終了時にデバッグ・モード無効
	}()

	phraseOriginal := "Hello, world!"
	phraseTranslated := "世界よ、こんにちは！"

	out := capturer.CaptureOutput(func() {
		// Set でキャッシュに登録
		_ = tmpCache.Set(phraseOriginal, phraseTranslated)

		// Get でキャッシュから取得
		result, _ := tmpCache.Get(phraseOriginal)

		expect := phraseTranslated
		actual := result

		assert.Equal(t, expect, actual)
	})

	assert.Contains(t, out, "badger", "on debug mode badger log should be printed")
	assert.Contains(t, out, "INFO", "on debug mode DB info should be printed")
}
