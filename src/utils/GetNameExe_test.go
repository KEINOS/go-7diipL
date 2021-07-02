package utils_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetNameExe(t *testing.T) {
	expect := "utils"
	actual := utils.GetNameExe()

	assert.Equal(t, expect, actual)
}
