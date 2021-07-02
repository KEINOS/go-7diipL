package cache

import "github.com/dgraph-io/badger/v3"

// NameDirCacheDefault はキャッシュ時のデフォルトのディレクトリ名もしくはプレフィックス（接頭辞）です.
//
// テンポラリ・ディレクトリの、この名前の付いたディレクトリは削除可能です.
const NameDirCacheDefault = "QiiTrans"

// TCache はキャッシュ DB の構造体です. [badger](https://github.com/dgraph-io/badger) のラッパーです.
type TCache struct {
	cacheDB      *badger.DB
	NameDirCache string
	PathDirTemp  string
}
