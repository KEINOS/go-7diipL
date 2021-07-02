package cache_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/cache"
	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
)

func TestClearAll(t *testing.T) {
	c := cache.New(t.Name())

	defer c.ClearAll()

	pathDirCache := c.GetPathDirCache()
	phraseOriginal := "foo bar"
	phraseTranslated := "hoge fuga"

	// Set cache
	_ = c.Set(phraseOriginal, phraseTranslated)

	// Get cache
	result, _ := c.Get(phraseOriginal)

	// Test save
	expect := phraseTranslated
	actual := result

	assert.Equal(t, expect, actual)
	assert.DirExists(t, pathDirCache, "before ClearAll() call, the cache dir should exist")

	// Run ClearAll
	c.ClearAll()

	assert.NoDirExists(t, pathDirCache, "after ClearAll() call, cache dir should not exist")

	// Run secondary call
	out := capturer.CaptureOutput(func() {
		c.ClearAll()
	})

	// Test delete
	assert.Empty(t, out, "if cache dir didn't exist then ClearAll() should retrun nothing")
}

func TestClearAll_on_opened_DB(t *testing.T) {
	c := cache.New(t.Name())

	defer c.ClearAll()

	out := capturer.CaptureOutput(func() {
		err := c.OpenDB()
		assert.NoError(t, err)

		c.ClearAll()
	})

	assert.Empty(t, out, "if the DB is opend it should close before deletion")
}
