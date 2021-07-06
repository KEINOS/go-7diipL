package app_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/app"
	"github.com/stretchr/testify/assert"
)

func TestGetMsgHelpUsage(t *testing.T) {
	nameAppDummy := "foo"
	nameExeDummy := "bar"

	out := app.GetMsgHelpUsage(nameAppDummy, nameExeDummy)

	assert.Contains(t, out, nameAppDummy, "it should parse the template with the arg value")
	assert.Contains(t, out, nameExeDummy, "it should parse the template with the arg value")
}

func TestGetMsgHelpUsage_missing_app_name(t *testing.T) {
	nameAppDummy := ""
	nameExeDummy := "bar"

	out := app.GetMsgHelpUsage(nameAppDummy, nameExeDummy)

	assert.Contains(t, out, app.NameDefault, "if 1st arg is empty then it should use the default app name")
}
