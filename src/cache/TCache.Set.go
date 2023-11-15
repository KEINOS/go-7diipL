package cache

import (
	"github.com/dgraph-io/badger/v3"
	"github.com/pkg/errors"
)

// Set は key に value の値を割り当ててキャッシュします.
// この関数は、主に翻訳前の文を key、翻訳後の文を value としてキャッシュに使われます.
//
// 引数 key は長さに依存しません（内部で 64 バイトの固定長にハッシュ化されるため）.
// 呼び出しごとにキャッシュ DB を open/close するため、大量の連続呼び出しには向きません.
// 大量に登録する場合は SetList() 関数を利用してください.
func (c *TCache) Set(key string, value string) error {
	err := c.OpenDB()
	if err != nil {
		return errors.Wrap(err, "failed to open cache db")
	}

	defer c.CloseDB()

	// key からハッシュ値を取得し DB のキー名として利用
	keyHashed := c.hash64(key)

	// 読み書きトランザクション
	err = c.cacheDB.Update(func(txn *badger.Txn) error {
		err = txn.SetEntry(badger.NewEntry(keyHashed, []byte(value)))

		return errors.Wrap(err, "failed to update cache")
	})

	return errors.Wrap(err, "failed to set cache")
}
