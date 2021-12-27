package app_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/app"
	"github.com/stretchr/testify/assert"
)

func TestSetEngine(t *testing.T) {
	nameEngine := "deepl"

	appTest := app.New("", t.Name())

	err := appTest.SetEngine(nameEngine)

	assert.NoError(t, err)

	defer appTest.Engine.Cache.ClearAll()

	expect := nameEngine
	actual := appTest.Engine.NameEngine

	assert.Equal(t, expect, actual)
}
