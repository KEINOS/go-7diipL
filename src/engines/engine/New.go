package engine

import "github.com/Qithub-BOT/QiiTrans/src/cache"

// New は翻訳エンジンの基本構造体の新規オブジェクトのポインタを返します.
func New(cacheID ...string) *Properties {
	tmpObj := new(Properties)

	tmpObj.Cache = cache.New(cacheID...)

	return tmpObj
}
