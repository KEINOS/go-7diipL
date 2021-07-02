package cache

import (
	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/dgraph-io/badger/v3"
)

// OpenDB はキャッシュ用の DB をオープンします. キャッシュ・ディレクトリが存在しない場合は、テンポラリディレクトリに新規に作成します.
//
// このメソッドは、Get() Set() でも呼び出されるため、利用時に OpenDB() CloseDB() を行う必要はありません.
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
