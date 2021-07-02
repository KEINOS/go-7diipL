package cache_test

import (
	"fmt"

	"github.com/Qithub-BOT/QiiTrans/src/cache"
)

func ExampleNew() {
	// 汎用のキャッシュにしたくない場合は ID を指定する。この場合は "my sample DB".
	c := cache.New("my sample DB")

	// 終了時にキャッシュを削除したい場合は defer で ClearAll する
	defer c.ClearAll()

	phraseOriginal := "Hello, world!"
	phraseTranslated := "世界よ、こんにちは！"

	// Set でキャッシュに登録
	_ = c.Set(phraseOriginal, phraseTranslated)

	// Get でキャッシュから取得
	result, _ := c.Get(phraseOriginal)

	fmt.Println(result)
	// Output: 世界よ、こんにちは！
}
