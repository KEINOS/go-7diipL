package utils_test

import (
	"errors"
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/stretchr/testify/assert"
)

func TestPanicOnErr(t *testing.T) {
	msg := "foobar"
	err := errors.New(msg)

	assert.PanicsWithError(t, msg, func() {
		utils.PanicOnErr(err)
	}, "if error is not a nil then it should panic with the given error")

	assert.NotPanics(t, func() {
		utils.PanicOnErr(nil)
	}, "if error is nil it should not panic")
}
