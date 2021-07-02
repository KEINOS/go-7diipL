package deepleng_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/engines/deepleng"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	ngin := deepleng.New()

	defer ngin.Cache.ClearAll()
	defer ngin.SetAPIKey("")

	ngin.Update = true

	inputText := "おはようございます。今日はとてもいい天気ですね。"

	outputText, isCache, err := ngin.Translate(inputText, "JA", "EN")

	expect := "Good morning. It's a very nice day today."
	actual := outputText

	assert.Nil(t, err)
	assert.False(t, isCache, "when Update property is true then isCache flag should return false")
	assert.Equal(t, expect, actual)
}
