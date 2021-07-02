package deepleng_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/engines/deepleng"
	"github.com/stretchr/testify/assert"
)

func TestIsAPIKeyFree(t *testing.T) {
	{
		sample := "foobarbuz:fx"
		result := deepleng.IsAPIKeyFree(sample)

		assert.True(t, result, "string ending with ':fx' should return true")
	}
	{
		sample := "foobarbuz"
		result := deepleng.IsAPIKeyFree(sample)

		assert.False(t, result, "string not ending with ':fx' should return false")
	}
}
