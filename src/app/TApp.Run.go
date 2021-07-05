package app

import (
	"errors"
	"fmt"

	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/mkideal/cli"
)

// ForceFail は Run メソッドを強制的に失敗させる（1 を返す）ためのフラグです.
// 主にテスト目的で使われます.
var ForceFail bool = false

// Run メソッドはアプリを実行します. 実行に成功した場合は nil を、失敗した場合は error を返します.
//
// テスト目的で強制的に失敗させる場合は ForceFail を true に設定します.
func (a *TApp) Run() int {
	if ForceFail {
		return utils.FAILURE
	}

	tmpFlag := new(TFlagOptions)
	tmpFlag.SetHelpMsg()

	return cli.Run(tmpFlag, a.run)
}

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

func (a *TApp) run(ctx *cli.Context) error {
	if err := a.SetArgValue(ctx); err != nil {
		return err
	}

	// バージョン情報表示
	if a.Argv.Version {
		fmt.Println(a.GetVersion())

		return nil
	}

	if err := a.PreRun(); err != nil {
		return err
	}

	// 翻訳する言語のリストを取得（フラグやオプション以外の引数）
	orderLang := ctx.Args()

	// 翻訳エンジン取得
	if err := a.SetEngine(a.Argv.NameEngine); err != nil {
		return err
	}

	// キャッシュを行うか
	a.Engine.Update = a.Argv.IsNoCache

	// APIキー/アクセス・トークンが引数で渡された場合、環境変数にセットする.
	// 既存の環境変数があった場合に備えて、リカバリ用の関数を取得する.
	deferFunc := a.Engine.SetAPIKey(a.Argv.APIKey)

	defer deferFunc() // 既存の環境変数にリカバリ

	// 実行前のキャッシュ削除
	if a.Argv.ClearBeforeRun {
		a.Engine.Cache.ClearAll()

		a.Argv.IsNoCache = true
	}

	// 標準入力から翻訳元のデータを取得
	input, err := utils.GetSTDIN()
	if err != nil {
		return err
	}

	// 翻訳の実行
	result, err := a.Translate(orderLang, input)
	if err != nil {
		return err
	}

	// 最終結果出力
	ctx.String("%s", result)

	return nil
}
