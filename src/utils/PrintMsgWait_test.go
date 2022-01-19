package utils_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
)

func TestPrintMsgWait(t *testing.T) {
	out := capturer.CaptureStdout(func() {
		setoff := utils.PrintMsgWait("please wait ...")
		setoff()
	})

	assert.Contains(t, out, "please wait ...\r")
	assert.Contains(t, out, "\r               \r")
}
