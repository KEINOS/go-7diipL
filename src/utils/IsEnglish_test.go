package utils_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/stretchr/testify/assert"
)

func TestIsEnglish(t *testing.T) {
	{
		sample := `A perennial also-ran, Stallings won his seat when longtime lawmaker David Holmes
		died 11 days after the filing deadline. Suddenly, Stallings was a shoo-in, not
		the long shot. In short order, the Legislature attempted to pass a law allowing
		former U.S. Rep. Carolyn Cheeks Kilpatrick to file; Stallings challenged the
		law in court and won. Kilpatrick mounted a write-in campaign, but Stallings won.`

		result := utils.IsEnglish(sample)
		assert.True(t, result, "string in English should return true")
	}
	{
		sample := "これは日本語です。明らかに日本語です。"

		result := utils.IsEnglish(sample)
		assert.False(t, result, "string in Japanese should return false")
	}
	{
		sample := "これは日本語です。\nしかし、英文 'This is a sample' も入っています。\n\tOr like this one."

		result := utils.IsEnglish(sample)
		assert.False(t, result, "string in Japanese with english mixed should return false")
	}
}
