package cache_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/cache"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTCache_Set(t *testing.T) {
	tmpCache := new(cache.TCache)

	err := tmpCache.Set("key", "value")

	require.Error(t, err,
		"it should error if cache DB is not initialized")
	assert.Contains(t, err.Error(), "failed to open cache db",
		"it should contain the error reason")
}
