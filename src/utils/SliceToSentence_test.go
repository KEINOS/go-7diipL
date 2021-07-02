package utils_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/stretchr/testify/assert"
)

func TestSliceSentences(t *testing.T) {
	sample := "これはサンプルです。This is a sample. Yet another sample."
	expect := []string{
		"これはサンプルです。",
		"This is a sample.",
		"Yet another sample.",
	}
	actual := utils.SliceSentences(sample)

	assert.Equal(t, expect, actual)
}

func TestSliceSentences_decimal(t *testing.T) {
	sample := "これはサンプルです。\nThis is a sample with dot such as 3.0. Yet another sample."
	expect := []string{
		"これはサンプルです。",
		"This is a sample with dot such as 3.0.",
		"Yet another sample.",
	}
	actual := utils.SliceSentences(sample)

	assert.Equal(t, expect, actual)
}

func TestSliceSentences_english(t *testing.T) {
	{
		sample := "Hi there. Does this really work?"
		expect := []string{
			"Hi there.",
			"Does this really work?",
		}
		actual := utils.SliceSentences(sample)

		assert.Equal(t, expect, actual)
	}
	{
		sample := `A perennial also-ran, Stallings won his seat when longtime lawmaker David Holmes
		died 11 days after the filing deadline. Suddenly, Stallings was a shoo-in, not
		the long shot. In short order, the Legislature attempted to pass a law allowing
		former U.S. Rep. Carolyn Cheeks Kilpatrick to file; Stallings challenged the
		law in court and won. Kilpatrick mounted a write-in campaign, but Stallings won.`
		expect := []string{
			"A perennial also-ran, Stallings won his seat when longtime lawmaker David Holmes died " +
				"11 days after the filing deadline.",
			"Suddenly, Stallings was a shoo-in, not the long shot.",
			"In short order, the Legislature attempted to pass a law allowing former U.S. Rep. " +
				"Carolyn Cheeks Kilpatrick to file; Stallings challenged the law in court and won.",
			"Kilpatrick mounted a write-in campaign, but Stallings won.",
		}
		actual := utils.SliceSentences(sample)

		assert.Equal(t, expect, actual)
	}
}
