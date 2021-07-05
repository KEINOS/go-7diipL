package app

import (
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

func (a *TApp) run(ctx *cli.Context) error {
	// フラグ・オプションの値をセット
	if err := a.SetArgValue(ctx); err != nil {
		return err
	}

	// フラグ・オプションに対する事前処理（デバッグ・モード設定など）
	if err := a.PreRun(); err != nil {
		return err
	}

	// バージョン情報表示
	if a.Argv.Version {
		fmt.Println(a.GetVersion())

		return nil
	}

	// 翻訳する言語のリストを取得（フラグやオプション以外の引数）
	orderLang := ctx.Args()

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

	// API 情報の表示
	if a.Argv.ShowInfo {
		quotaLeft, err := a.Engine.GetQuotaLeft()
		if err != nil {
			return err
		}

		utils.EchoSTDERR("\n[INFO]: 残り文字数: %v\n", utils.DelimitComma(quotaLeft))
	}

	return nil
}
