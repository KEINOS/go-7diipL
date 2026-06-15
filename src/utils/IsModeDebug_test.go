package utils_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/stretchr/testify/assert"
)

//nolint:paralleltest // due to the monkey patching of global variable(s)
func TestIsModeDebug(t *testing.T) {
	utils.SetModeDebug(true)

	defer utils.SetModeDebug(false)

	result := utils.IsModeDebug()

	assert.True(t, result, "on debug mode it should return ture")
}
