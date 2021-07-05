package utils_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/stretchr/testify/assert"
)

func TestDelimitComma(t *testing.T) {
	input := 1000000

	expect := "1,000,000"
	actual := utils.DelimitComma(input)

	assert.Equal(t, expect, actual, "it should be a comma-delimited string")
}
