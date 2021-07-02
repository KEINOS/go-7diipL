package utils_test

import (
	"fmt"
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
)

func TestLogDebug(t *testing.T) {
	defer utils.SetModeDebug(false)

	msg := "foobar"

	{
		utils.SetModeDebug(false)

		out := capturer.CaptureOutput(func() {
			utils.LogDebug(msg)
		})

		assert.Empty(t, out, "if debug mode is false then it should output nothing")
	}
	{
		utils.SetModeDebug(true)

		expect := fmt.Sprintf("[LOG]: %s\n", msg)
		actual := capturer.CaptureStderr(func() {
			utils.LogDebug(msg)
		})

		assert.Equal(t, expect, actual, "single arg should be treated as a string")
	}
	{
		utils.SetModeDebug(true)

		expect := fmt.Sprintf("[LOG]: %s\n", msg)
		actual := capturer.CaptureStderr(func() {
			utils.LogDebug("%s", msg)
		})

		assert.Equal(t, expect, actual, "more than one arg should be treated as format style")
	}
}
