package app_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/app"
	"github.com/Qithub-BOT/QiiTrans/src/engines/deepleng"
	"github.com/stretchr/testify/assert"
)

func TestNewEngine(t *testing.T) {
	expect := deepleng.New(t.Name())

	defer expect.Cache.ClearAll() // 終了後に削除

	appTest := app.New(t.Name())
	actual, err := appTest.NewEngine("deepl", t.Name())
	assert.NoError(t, err)

	defer actual.Cache.ClearAll() // 終了後に削除

	assert.IsType(t, expect, actual, "it should be the same object type")
	assert.Equal(t, expect.NameEngine, actual.NameEngine, "it should be the same engine type")
}
