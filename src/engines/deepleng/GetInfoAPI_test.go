package deepleng_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/engines/deepleng"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetInfoAPI(t *testing.T) {
	ngin := deepleng.New()

	defer ngin.Cache.ClearAll()

	quotaLeft, err := ngin.GetQuotaLeft()

	require.NoError(t, err)
	t.Logf("Current API quota left: %v characters", quotaLeft)

	expect := 0
	actual := quotaLeft

	assert.NotEqual(t, expect, actual, "no quota left. seems that API limit exceed or wrong return")
	assert.Greater(t, actual, expect, "it should return the char number left")
}

func TestGetInfoAPI_no_token_set(t *testing.T) {
	ngin := deepleng.New()

	defer ngin.Cache.ClearAll()

	ngin.NameVarEnvAPIKey = "UNKNOWN_KEY_IN_ENV"

	_, err := ngin.GetQuotaLeft()

	require.Error(t, err, "if required env key not found then it should return an error")
}
