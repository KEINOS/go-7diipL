package deepleng_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/engines/deepleng"
	"github.com/stretchr/testify/assert"
)

func TestGetInfoAPI(t *testing.T) {
	ngin := deepleng.New()

	defer ngin.Cache.ClearAll()

	quotaLeft, err := ngin.GetQuotaLeft()

	assert.Nil(t, err)

	t.Logf("Current API quota left: %v characters", quotaLeft)

	assert.NotEqual(t, quotaLeft, 0, "no quota left. seems that API limit exceed or wrong return")
	assert.Greater(t, quotaLeft, 0, "it should return the char number left")
}

func TestGetInfoAPI_no_token_set(t *testing.T) {
	ngin := deepleng.New()

	defer ngin.Cache.ClearAll()

	ngin.NameVarEnvAPIKey = "UNKNOWN_KEY_IN_ENV"

	_, err := ngin.GetQuotaLeft()

	assert.Error(t, err, "if required env key not found then it should return an error")
}
