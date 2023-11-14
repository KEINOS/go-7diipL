package utils_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/helperfunc"
	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetSTDIN(t *testing.T) {
	const expect = "foobar"

	utils.ValSTDINDummy = expect

	defer func() {
		utils.ValSTDINDummy = ""
	}()

	actual, err := utils.GetSTDIN()

	assert.NoError(t, err)
	assert.Equal(t, expect, actual)
}

func TestGetSTDIN_mock_stdin(t *testing.T) {
	userInput := "foo bar buzz"

	funcDefer := helperfunc.MockSTDIN(t, userInput)
	defer funcDefer() // clean up

	expect := userInput
	actual, err := utils.GetSTDIN()

	assert.NoError(t, err)
	assert.Equal(t, expect, actual)
}

func TestGetSTDIN_forced_error(t *testing.T) {
	userInput := "foo bar buzz"

	utils.ForceErrorGetSTDIN = true
	defer func() { utils.ForceErrorGetSTDIN = true }()

	funcDefer := helperfunc.MockSTDIN(t, userInput)
	defer funcDefer() // clean up

	out, err := utils.GetSTDIN()

	assert.Error(t, err)
	assert.Empty(t, out, "on error it should be empty")
}
