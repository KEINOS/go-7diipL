package cache_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/cache"
	"github.com/stretchr/testify/assert"
)

func TestGetPathDirCache(t *testing.T) {
	c := cache.New(t.Name())

	defer c.ClearAll()

	expect := filepath.Join(os.TempDir(), "QiiTrans") + "_" + t.Name()
	actual := c.GetPathDirCache()

	assert.Equal(t, expect, actual, "it should return the path to the cached file")
}
