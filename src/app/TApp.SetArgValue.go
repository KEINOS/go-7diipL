package app

import (
	"errors"

	"github.com/mkideal/cli"
)

func (a *TApp) SetArgValue(ctx *cli.Context) error {
	// フラグやオプション引数を取得
	argv, ok := ctx.Argv().(*TFlagOptions)
	if !ok {
		return errors.New("コマンド引数の読み込みに失敗しました")
	}

	// 取得したフラグやオプションを割り当て
	a.Argv = argv

	// ヘルプの再設定
	a.Argv.SetHelpMsg()

	return nil
}
