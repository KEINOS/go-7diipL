package engine

import (
	"fmt"
	"os"

	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"golang.org/x/xerrors"
)

// SetAPIKey はコマンド引数から取得したアクセス・トークン "apikey" を翻訳エンジンが使えるようにセットします.
func (p *Properties) SetAPIKey(apikey string) func() {
	if p.NameVarEnvAPIKey == "" {
		msg := "翻訳エンジンが参照する API キー/アクセス・トークンの環境変数名が指定されていません"
		fmt.Fprintln(os.Stderr, xerrors.New(msg))
	}

	p.apiKey = apikey

	deferFunc := func() {}

	if p.apiKey == "" {
		return deferFunc
	}

	// 既存の同名の環境変数がないかチェック
	oldKey := os.Getenv(p.NameVarEnvAPIKey)

	err := os.Setenv(p.NameVarEnvAPIKey, p.apiKey)
	utils.ExitOnErr(err)

	// 呼び出し元の Defer 用に環境変数を元に戻す関数を返す
	return func() {
		utils.ExitOnErr(os.Setenv(p.NameVarEnvAPIKey, oldKey))
	}
}
