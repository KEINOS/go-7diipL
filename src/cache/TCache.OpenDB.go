package cache

import (
	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/dgraph-io/badger/v3"
)

func (c *TCache) OpenDB() (err error) {
	// キャッシュ用の DB をオープン。存在しない場合は新規作成してオープン
	pathDirCache := c.GetPathDirCache()
	optionsDB := badger.DefaultOptions(pathDirCache)

	// デフォルトで INFO 情報が表示されるため、デバッグ・モード時以外は非表示
	if !utils.IsModeDebug() {
		optionsDB.Logger = nil
	}

	c.cacheDB, err = badger.Open(optionsDB)

	return err
}
