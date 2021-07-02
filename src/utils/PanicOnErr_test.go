package utils_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/stretchr/testify/assert"
	"golang.org/x/xerrors"
)

func TestPanicOnErr(t *testing.T) {
	msg := "foobar"
	err := xerrors.New(msg)

	assert.PanicsWithError(t, msg, func() {
		utils.PanicOnErr(err)
	}, "if error is not a nil then it should panic with the given error")

	assert.NotPanics(t, func() {
		utils.PanicOnErr(nil)
	}, "if error is nil it should not panic")
}
