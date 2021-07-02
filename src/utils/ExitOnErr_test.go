package utils_test

import (
	"errors"
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
)

func TestExitOnErr(t *testing.T) {
	oldOsExit := utils.OsExit

	defer func() {
		utils.OsExit = oldOsExit
	}()

	msg := "foobar"
	err := errors.New(msg)

	var received int

	utils.OsExit = func(code int) {
		received = code
	}

	out := capturer.CaptureStderr(func() {
		utils.ExitOnErr(err)
	})

	assert.Contains(t, out, msg, "the output message should be to stderr")

	expect := received
	actual := utils.FAILURE

	assert.Equal(t, expect, actual, "it should exit with code %v", utils.FAILURE)
}
