package cache_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/cache"
	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClearAll(t *testing.T) {
	tmpCache := cache.New(t.Name())

	defer tmpCache.ClearAll()

	pathDirCache := tmpCache.GetPathDirCache()
	phraseOriginal := "foo bar"
	phraseTranslated := "hoge fuga"

	// Set cache
	err := tmpCache.Set(phraseOriginal, phraseTranslated)
	require.NoError(t, err, "failed to set cache during test preparation")

	// Get cache
	result, err := tmpCache.Get(phraseOriginal)
	require.NoError(t, err, "failed to get cache during test preparation")

	// Test save
	expect := phraseTranslated
	actual := result

	assert.Equal(t, expect, actual)
	assert.DirExists(t, pathDirCache, "before ClearAll() call, the cache dir should exist")

	// Run ClearAll
	tmpCache.ClearAll()

	assert.NoDirExists(t, pathDirCache, "after ClearAll() call, cache dir should not exist")

	// Run secondary call
	out := capturer.CaptureOutput(func() {
		tmpCache.ClearAll()
	})

	// Test delete
	assert.Empty(t, out, "if cache dir didn't exist then ClearAll() should retrun nothing")
}

func TestClearAll_on_opened_DB(t *testing.T) {
	tmpCache := cache.New(t.Name())

	defer tmpCache.ClearAll()

	out := capturer.CaptureOutput(func() {
		err := tmpCache.OpenDB()
		require.NoError(t, err)

		tmpCache.ClearAll()
	})

	assert.Empty(t, out, "if the DB is opend it should close before deletion")
}
