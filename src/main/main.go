package main

import (
	"errors"
	"fmt"

	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/mkideal/cli"
)

// versionApp はアプリのバージョン情報です. デフォルト値以外にセットする場合は、アプリのバイナリをビルドする際に設定してください.
var versionApp string = VersionAppDefault

func main() {
	result := cli.Run(new(TFlag), func(ctx *cli.Context) error {
		/*
			フラグやオプションのバリデーションは TFlag.Validate() で行っています.
		*/

		// フラグやオプション引数を取得
		argv, ok := ctx.Argv().(*TFlag)
		if !ok {
			return errors.New("コマンド引数の読み込みに失敗しました")
		}

		// バージョン情報表示
		if argv.Version {
			fmt.Println(GetVersionApp())

			return nil
		}

		// アプリの初期化
		if err := PreRun(argv); err != nil {
			return err
		}

		// 翻訳する言語のリストを、フラグやオプション以外の引数から取得
		orderLang := ctx.Args()

		// 翻訳エンジン取得
		ngin, err := NewEngine(argv.NameEngine)
		if err != nil {
			return err
		}

		// キャッシュを行うか
		ngin.Update = argv.IsNoCache

		// APIキー/アクセス・トークンが引数で渡された場合、環境変数にセットする.
		// 既存の環境変数があった場合に備えて、リカバリ用の関数を取得する.
		deferFunc := ngin.SetAPIKey(argv.APIKey)

		defer deferFunc() // 既存の環境変数にリカバリ

		// 実行前のキャッシュ削除
		if argv.ClearBeforeRun {
			ngin.Cache.ClearAll()

			argv.IsNoCache = true
		}

		// 翻訳実行
		return Run(ngin, orderLang, ctx)
	})

	utils.OsExit(result)
}
