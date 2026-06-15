package cache_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/cache"
	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//nolint:paralleltest // due to the monkey patching of global variable(s)
func TestGet_fail_to_openDB(t *testing.T) {
	tmpCache := new(cache.TCache)

	_, err := tmpCache.Get("foo bar")

	require.Error(t, err, "if DB fails to open it should return an error")
}

//nolint:paralleltest // due to the monkey patching of global variable(s)
func TestGet_not_cached(t *testing.T) {
	tmpCache := cache.New(t.Name())

	defer tmpCache.ClearAll()

	phraseOriginal := "unset(uncached) phrase"

	// Get cache
	out := capturer.CaptureOutput(func() {
		_, err := tmpCache.Get(phraseOriginal)

		require.Error(t, err, "uncached phrase should return an error")
	})

	assert.Empty(t, out, "on Get() fail, it should not printout anything")
}
