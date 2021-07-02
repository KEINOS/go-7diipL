package cache

import "github.com/dgraph-io/badger/v3"

// Set は value の値を key に割り当ててキャッシュします.
//
// 引数 key は、内部で 64 バイトの固定長にハッシュ化されるため、長さに依存しません.
// この関数は、主に翻訳前の文と、翻訳後の文のキャッシュに使われます.
// 呼び出しごとにキャッシュ DB を open/close するため、大量の連続呼び出しには向きません.
// 大量に登録する場合は SetList() 関数を利用してください.
func (c *TCache) Set(key string, value string) (err error) {
	if err = c.OpenDB(); err == nil {
		defer c.CloseDB()

		// original から DB の key となるハッシュ値を取得
		key := c.hash64(key)

		// 読み書きトランザクション
		err = c.cacheDB.Update(func(txn *badger.Txn) error {
			return txn.SetEntry(badger.NewEntry(key, []byte(value)))
		})
	}

	return err
}
