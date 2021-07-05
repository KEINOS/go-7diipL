package app

import (
	"errors"

	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/mkideal/cli"
)

// Validate メソッドは、cli.Validator インターフェースの実装です.
//
// コマンドのフラグ、オプションや引数のバリデーションを行います. cli.Run() で指定された関数が実行される前に呼び出されます.
func (argv *TFlagOptions) Validate(ctx *cli.Context) error {
	// オプションやグラグ以外の引数を取得
	inputArgs := ctx.Args()

	// バリデーションを無視するフラグ
	switch {
	case argv.Version:
		return nil
	}

	// 必須引数の個数
	switch len(inputArgs) {
	case 0:
		return errors.New("引数が足りません。最低でも翻訳元と翻訳先の言語を指定してください")
	case 1:
		return errors.New("引数が足りません。翻訳先の言語を指定してください")
	}

	// 引数の言語指定のバリデーション
	for _, lang := range inputArgs {
		if !utils.IsValidLang(lang) {
			return errors.New("Bad lang. 引数の言語名が未定義のものです: " + lang)
		}
	}

	return nil
}
