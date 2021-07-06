package app_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/app"
	"github.com/kami-zh/go-capturer"
	"github.com/mkideal/cli"
	"github.com/stretchr/testify/assert"
)

func TestSetArgValue(t *testing.T) {
	appTest := app.New(t.Name())
	argDummy := new(cli.Context)

	out := capturer.CaptureOutput(func() {
		err := appTest.SetArgValue(argDummy)

		assert.Error(t, err)
	})

	assert.Empty(t, out)
}
