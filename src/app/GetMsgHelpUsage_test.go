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
