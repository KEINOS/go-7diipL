package deepleng_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/engines/deepleng"
	"github.com/Qithub-BOT/QiiTrans/src/engines/engine"
	"github.com/stretchr/testify/assert"
)

func TestGetURLBaseAPI(t *testing.T) {
	e := engine.New(t.Name())

	defer e.Cache.ClearAll()

	{
		e.IsAccountFree = true

		expect := "https://api-free.deepl.com"
		actual := deepleng.GetURLBaseAPI(e)

		assert.Equal(t, expect, actual)
	}

	{
		e.IsAccountFree = false

		expect := "https://api.deepl.com"
		actual := deepleng.GetURLBaseAPI(e)

		assert.Equal(t, expect, actual)
	}
}
