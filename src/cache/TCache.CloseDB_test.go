package cache_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/cache"
	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
)

func TestCloseDB(t *testing.T) {
	c := cache.New(t.Name())

	defer c.ClearAll()

	err := c.OpenDB()

	assert.NoError(t, err)
	assert.NotPanics(t, func() {
		c.CloseDB()
	}, "closing an opened DB should not panic")
}

func TestCloseDB_debug_mode(t *testing.T) {
	c := cache.New(t.Name())

	utils.SetModeDebug(true)

	defer func() {
		utils.SetModeDebug(false)
		c.ClearAll()
	}()

	out := capturer.CaptureOutput(func() {
		c.CloseDB()
	})

	contain := "not opened"

	assert.Contains(t, out, contain, "closing a not opened DB should log on debug mode")
}
