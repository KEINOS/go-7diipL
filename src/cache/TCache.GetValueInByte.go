package cache

import "github.com/dgraph-io/badger/v3"

// GetValueInByte はバイトデータにハッシュ化された keyHashed キーを使ってキャッシュを探索しバイトデータで返します.
func (c *TCache) GetValueInByte(keyHashed []byte) ([]byte, error) {
	var valCopy []byte

	err := c.cacheDB.View(func(txn *badger.Txn) error {
		// 念のための初期化
		valCopy = []byte("")

		// キーを指定してアイテム取得
		item, err := txn.Get(keyHashed)
		if err != nil {
			return err
		}

		// アイテムの byte データ取得
		err = item.Value(func(val []byte) error {
			valCopy = append([]byte{}, val...)

			return nil
		})

		return err
	})

	return valCopy, err
}
