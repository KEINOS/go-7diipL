package main

import (
	"github.com/Qithub-BOT/QiiTrans/src/engines/engine"
	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/mkideal/cli"
)

// Run は、翻訳エンジンのインスタンスを使って翻訳を実行し、結果をコンテキストに渡します.
func Run(ngin *engine.Properties, orderLang []string, ctx *cli.Context) error {
	// 標準入力から翻訳元のデータを取得
	input, err := utils.GetSTDIN()
	if err != nil {
		return err
	}

	// 翻訳の実行
	result, err := Translate(ngin, orderLang, input)
	if err != nil {
		return err
	}

	// 最終結果出力
	ctx.String("%s", result)

	return nil
}
