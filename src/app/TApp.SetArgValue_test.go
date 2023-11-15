package app_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/app"
	"github.com/kami-zh/go-capturer"
	"github.com/mkideal/cli"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSetArgValue(t *testing.T) {
	appTest := app.New("", t.Name())
	argDummy := new(cli.Context)

	out := capturer.CaptureOutput(func() {
		err := appTest.SetArgValue(argDummy)

		require.Error(t, err, "missing engine name should return an error")
	})

	assert.Empty(t, out, "on error it should not print anything")
}
