package deepleng_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/engines/deepleng"
	"github.com/Qithub-BOT/QiiTrans/src/engines/engine"
	"github.com/stretchr/testify/assert"
)

func TestGetURLBaseAPI(t *testing.T) {
	eng := engine.New(t.Name())

	defer eng.Cache.ClearAll()

	{
		eng.IsAccountFree = true

		expect := "https://api-free.deepl.com"
		actual := deepleng.GetURLBaseAPI(eng)

		assert.Equal(t, expect, actual)
	}

	{
		eng.IsAccountFree = false

		expect := "https://api.deepl.com"
		actual := deepleng.GetURLBaseAPI(eng)

		assert.Equal(t, expect, actual)
	}
}
