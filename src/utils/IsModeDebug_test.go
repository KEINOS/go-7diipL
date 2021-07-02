package utils_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/stretchr/testify/assert"
)

func TestIsModeDebug(t *testing.T) {
	utils.SetModeDebug(true)

	defer utils.SetModeDebug(false)

	result := utils.IsModeDebug()

	assert.True(t, result, "on debug mode it should return ture")
}
