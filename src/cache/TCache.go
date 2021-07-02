package cache

import "github.com/dgraph-io/badger/v3"

const (
	NameDirCacheDefault = "QiiTrans"
)

type TCache struct {
	cacheDB      *badger.DB
	NameDirCache string
	PathDirTemp  string
}
