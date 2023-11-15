package cache_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/cache"
	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCloseDB(t *testing.T) {
	tmpCache := cache.New(t.Name())

	defer tmpCache.ClearAll()

	err := tmpCache.OpenDB()

	require.NoError(t, err)
	assert.NotPanics(t, func() {
		tmpCache.CloseDB()
	}, "closing an opened DB should not panic")
}

func TestCloseDB_debug_mode(t *testing.T) {
	tmpCache := cache.New(t.Name())

	utils.SetModeDebug(true)

	defer func() {
		utils.SetModeDebug(false)
		tmpCache.ClearAll()
	}()

	out := capturer.CaptureOutput(func() {
		tmpCache.CloseDB()
	})

	contain := "not opened"

	assert.Contains(t, out, contain, "closing a not opened DB should log on debug mode")
}
