package cache

import (
	"os"

	"github.com/Qithub-BOT/QiiTrans/src/utils"
)

// ClearAll はキャッシュ・データを削除します.
func (c *TCache) ClearAll() {
	pathDirCache := c.GetPathDirCache()

	if !utils.IsDir(pathDirCache) {
		return
	}

	if c.cacheDB != nil && !c.cacheDB.IsClosed() {
		c.CloseDB()
	}

	err := os.RemoveAll(pathDirCache)

	utils.PanicOnErr(err)
}
