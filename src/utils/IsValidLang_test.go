package utils_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/stretchr/testify/assert"
)

func TestIsValidLang(t *testing.T) {
	t.Parallel()

	// Data Provider
	testCases := map[string]bool{
		"en":       true,
		"EN":       true,
		"english":  true,
		"ja":       true,
		"JA":       true,
		"JAPANESE": true,
		"unknown":  false,
	}

	for input, expect := range testCases {
		actual := utils.IsValidLang(input)

		assert.Equal(t, expect, actual, "Failed to detect lang: %s", input)
	}
}
