package cache

import (
	"github.com/Qithub-BOT/QiiTrans/src/utils"
)

// CloseDB はキャッシュの DB をクローズします. DB がオープンされていない場合は何もしません.
// クローズに失敗した場合はパニックを起こします.
//
// このメソッドは、Get() Set() でも呼び出されるため、利用時に OpenDB() CloseDB() を行う必要はありません.
func (c *TCache) CloseDB() {
	if c.cacheDB == nil {
		utils.LogDebug("closing DB that was not opened")

		return
	}

	err := c.cacheDB.Close()

	utils.PanicOnErr(err)
}
