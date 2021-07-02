package cache_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/cache"
	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
)

func TestGet_fail_to_openDB(t *testing.T) {
	c := new(cache.TCache)

	_, err := c.Get("foo bar")

	assert.Error(t, err, "if DB fails to open it should return an error")
}

func TestGet_not_cached(t *testing.T) {
	c := cache.New(t.Name())

	defer c.ClearAll()

	phraseOriginal := "unset(uncached) phrase"

	// Get cache
	out := capturer.CaptureOutput(func() {
		_, err := c.Get(phraseOriginal)
		assert.Error(t, err, "uncached phrase should return an error")
	})

	assert.Empty(t, out, "on Get() fail, it should not printout anything")
}
