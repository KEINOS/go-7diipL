package cache

import (
	"os"
	"strings"
)

// New は TCache の新規オブジェクトのポインタを返します.
// デフォルトで共通のキャッシュを利用します。各々異なるキャッシュにしたい場合は、引数に ID を指定します.
func New(cacheID ...string) *TCache {
	tmpObj := new(TCache)

	tmpObj.PathDirTemp = os.TempDir()
	tmpObj.NameDirCache = NameDirCacheDefault

	if len(cacheID) > 0 {
		// 引数や引数のホワイトスペース区切り文字列を "_" で結合する
		key := strings.Join(strings.Fields(strings.Join(cacheID, " ")), "_")

		tmpObj.NameDirCache = NameDirCacheDefault + "_" + key
	}

	return tmpObj
}
