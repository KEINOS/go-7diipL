package cache

import (
	"github.com/dgraph-io/badger/v3"
	"github.com/pkg/errors"
)

// GetValueInByte はバイトデータにハッシュ化された keyHashed キーを使ってキャッシュを探索しバイトデータで返します.
func (c *TCache) GetValueInByte(keyHashed []byte) ([]byte, error) {
	var valCopy []byte

	err := c.cacheDB.View(func(txn *badger.Txn) error {
		// 念のための初期化
		valCopy = []byte("")

		// キーを指定してアイテム取得
		item, err := txn.Get(keyHashed)
		if err != nil {
			return errors.Wrap(err, "failed to get item from cache")
		}

		// アイテムの byte データ取得
		err = item.Value(func(val []byte) error {
			valCopy = append([]byte{}, val...)

			return nil
		})

		return errors.Wrap(err, "failed to get value from cache")
	})

	return valCopy, errors.Wrap(err, "failed execute view transaction to read cache")
}
