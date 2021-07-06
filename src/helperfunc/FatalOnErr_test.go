package helperfunc_test

import (
	"errors"
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/helperfunc"
	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
)

func TestFatalOnErr(t *testing.T) {
	helperfunc.IsErrorDummy = true
	defer func() {
		helperfunc.IsErrorDummy = false
	}()

	msgErr := "foo bar"
	err := errors.New(msgErr)

	out := capturer.CaptureStderr(func() {
		helperfunc.FatalOnErr(t, err)
	})

	expect := "fatal error during test.\nErrMsg: foo bar"
	actual := out

	assert.Equal(t, expect, actual)
}

func TestFatalOnErr_additional_comment(t *testing.T) {
	helperfunc.IsErrorDummy = true
	defer func() {
		helperfunc.IsErrorDummy = false
	}()

	msgErr := "foo bar"
	err := errors.New(msgErr)

	comment1 := "this is an additional comment1"
	comment2 := "this is an additional comment2"

	out := capturer.CaptureStderr(func() {
		helperfunc.FatalOnErr(t, err, comment1, comment2)
	})

	expect := "fatal error during test.\n" +
		"ErrMsg: foo bar\n" +
		"this is an additional comment1\n" +
		"this is an additional comment2"
	actual := out

	assert.Equal(t, expect, actual)
}
