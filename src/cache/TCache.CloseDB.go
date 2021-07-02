package cache

import (
	"github.com/Qithub-BOT/QiiTrans/src/utils"
)

func (c *TCache) CloseDB() {
	if c.cacheDB == nil {
		utils.LogDebug("closing DB that was not opened")

		return
	}

	err := c.cacheDB.Close()

	utils.PanicOnErr(err)
}
