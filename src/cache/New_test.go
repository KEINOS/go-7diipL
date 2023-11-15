package cache_test

import (
	"fmt"

	"github.com/Qithub-BOT/QiiTrans/src/cache"
)

func ExampleNew() {
	// 汎用のキャッシュにしたくない場合は ID を指定する。この場合は "my sample DB".
	tmpCache := cache.New("my sample DB")

	// 終了時にキャッシュを削除したい場合は defer で ClearAll する
	defer tmpCache.ClearAll()

	phraseOriginal := "Hello, world!"
	phraseTranslated := "世界よ、こんにちは！"

	// Set でキャッシュに登録
	err := tmpCache.Set(phraseOriginal, phraseTranslated)
	if err != nil {
		panic(err)
	}

	// Get でキャッシュから取得
	result, err := tmpCache.Get(phraseOriginal)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
	// Output: 世界よ、こんにちは！
}
