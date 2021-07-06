package utils_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
)

func TestEchoSTDERR(t *testing.T) {
	input := "foo bar"
	expect := "Hoge foo bar Fuga"

	// 標準エラー出力チェック
	actual := capturer.CaptureStderr(func() {
		utils.EchoSTDERR("Hoge %s Fuga", input)
	})
	assert.Equal(t, expect, actual, "it should interpolate as fmt.Fprintf")

	// 標準出力チェック
	out := capturer.CaptureStdout(func() {
		utils.EchoSTDERR("Hoge %s Fuga", input)
	})
	assert.Empty(t, out, "it should not print out to STDOUT")
}
