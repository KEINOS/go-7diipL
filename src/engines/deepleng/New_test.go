package deepleng_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/engines/deepleng"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	ngin := deepleng.New()

	defer ngin.Cache.ClearAll()
	defer ngin.SetAPIKey("")

	ngin.Update = true

	inputText := "おはようございます。今日はとてもいい天気ですね。"

	outputText, isCache, err := ngin.Translate(inputText, "JA", "EN")

	expectList := []string{
		"Good morning. It's a very nice day today.",
		"Good morning. It is a very nice day today.",
	}

	actual := outputText

	require.NoError(t, err)
	assert.False(t, isCache, "when Update property is true then isCache flag should return false")
	assert.Contains(t, expectList, actual)
}
