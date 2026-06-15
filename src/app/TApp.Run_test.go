package app_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/app"
	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/stretchr/testify/assert"
)

//nolint:paralleltest // due to the monkey patching of global variable(s)
func TestRun_force_fail(t *testing.T) {
	appTest := app.New("", t.Name())

	app.ForceFailRun = true
	defer func() { app.ForceFailRun = false }()

	status := appTest.Run()

	expect := utils.FAILURE
	actual := status

	assert.Equal(t, expect, actual, "forcing fail should be non-zero")
}
