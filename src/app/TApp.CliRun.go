package app

import (
	"github.com/mkideal/cli"
)

// CliRun は app.Run() の本体です。cli.Run に登録して呼び出されるメソッドです.
func (a *TApp) CliRun(ctx *cli.Context) error {
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
		ctx.String("%s\n", a.GetVersion())

		return nil
	}

	// 翻訳する言語のリストを取得（フラグやオプション以外の引数）
	orderLang := ctx.Args()

	// APIキー/アクセス・トークンが引数で渡された場合、環境変数にセットする.
	// 既存の環境変数があった場合に備えて、リカバリ用の関数を取得する.
	deferFunc := a.Engine.SetAPIKey(a.Argv.APIKey)
	defer deferFunc() // 既存の環境変数にリカバリ

	// API 情報のみの表示
	if a.Argv.ShowInfoOnly {
		return a.PrintInfo(ctx)
	}

	// 実行前のキャッシュ削除
	if a.Argv.ClearBeforeRun {
		a.Engine.Cache.ClearAll()

		a.Argv.IsNoCache = true
	}

	if !a.Argv.IsPiped {
		return a.InteractiveTranslation(orderLang)
	}

	result, err := a.SingleShotTranslation(orderLang)
	if err != nil {
		return err
	}

	// 最終結果出力
	ctx.String("%s", result)

	// API 情報の出力
	if a.Argv.ShowInfo {
		return a.PrintInfo(ctx)
	}

	return nil
}
