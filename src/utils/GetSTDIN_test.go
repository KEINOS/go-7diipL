package utils_test

import (
	"os"
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetSTDIN(t *testing.T) {
	expect := "foobar"

	utils.ValSTDINDummy = expect

	defer func() {
		utils.ValSTDINDummy = ""
	}()

	actual, err := utils.GetSTDIN()

	assert.NoError(t, err)
	assert.Equal(t, expect, actual)
}

func TestGetSTDIN_mock_stdin(t *testing.T) {
	expect := "foo bar buzz"

	tmpFile, funcDefer := mockSTDIN(t, expect)

	defer funcDefer() // clean up

	oldStdin := os.Stdin

	defer func() { os.Stdin = oldStdin }() // Restore original Stdin

	// Mock stdin
	os.Stdin = tmpFile

	actual, err := utils.GetSTDIN()

	assert.NoError(t, err)
	assert.Equal(t, expect, actual)
}
