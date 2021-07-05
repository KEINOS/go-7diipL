package app_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/app"
	"github.com/stretchr/testify/assert"
)

func TestGetVersion(t *testing.T) {
	appTest := app.New()

	{
		expect := "QiiTrans dev version"
		actual := appTest.GetVersion()

		assert.Equal(t, expect, actual, "it should be dev version by default")
	}
	{
		appTest.Version = ""

		expect := "QiiTrans dev version"
		actual := appTest.GetVersion()

		assert.Equal(t, expect, actual, "missing version should be as dev")
	}
	{
		appTest.Version = "1.0.0\n"

		expect := "QiiTrans v1.0.0"
		actual := appTest.GetVersion()

		assert.Equal(t, expect, actual, "the 'v' should be added and no line break")
	}
}
