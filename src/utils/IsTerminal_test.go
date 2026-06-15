package utils_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/stretchr/testify/assert"
)

//nolint:paralleltest // due to the monkey patching of global variable(s)
func TestIsTerminal(t *testing.T) {
	result := utils.IsTerminal()

	assert.False(t, result, "calling from test should not be true")
}

//nolint:paralleltest // due to the monkey patching of global variable(s)
func TestIsTerminal_mock_result(t *testing.T) {
	utils.IsTerminalDummy = true

	defer func() { utils.IsTerminalDummy = false }()

	result := utils.IsTerminal()

	assert.True(t, result, "if IsTerminalDummy was set to true then it should return true as well")
}
