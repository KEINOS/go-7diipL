package engine

import (
	"fmt"
	"os"

	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"golang.org/x/xerrors"
)

// SetAPIKey はコマンド引数から取得したアクセストークン／認証キー（"apiKey"）を翻訳エンジンが使えるようにセットします.
//
// このメソッドは呼び出し元の defer 用に関数を返します。各々の翻訳エンジンが参照する環境変数に apiKey の値をセットするため、
// 既存の値があった場合は処理後 defer で元に戻せるようにするための関数です.
//
//	myEngine := deepleng.New("myCacheID")
//	myAPIKey := "foobar"
//	defer myEngine.SetAPIKey(myAPIKey)
func (p *Properties) SetAPIKey(apiKey string) func() {
	if p.NameVarEnvAPIKey == "" {
		msg := "翻訳エンジンが参照する API キー/アクセス・トークンの環境変数名が指定されていません"
		fmt.Fprintln(os.Stderr, xerrors.New(msg))
	}

	// 引数から API が指定されていない（空の）場合は、アプリのデフォルトの環境変数の値をセット
	if apiKey == "" {
		apiKey = os.Getenv(NameVarEnvAPIKeyDefault)
	}

	// アプリのデフォルトの環境変数に値が設定されていない（空の）場合は、翻訳 API の環境変数の値をセット
	if apiKey == "" {
		apiKey = os.Getenv(p.NameVarEnvAPIKey)
	}

	// エンジンのフィールドにアクセストークン（認証キー）をセット
	p.apiKey = apiKey

	// 翻訳 API の環境変数にセットするため、既存の環境変数の値をバックアップ
	oldKey := os.Getenv(p.NameVarEnvAPIKey)

	// 翻訳 API の環境変数にアクセストークン（認証キー）をセット
	err := os.Setenv(p.NameVarEnvAPIKey, p.apiKey)
	utils.ExitOnErr(err)

	// 呼び出し元の Defer 用に環境変数を元に戻す関数を返す
	return func() {
		utils.ExitOnErr(os.Setenv(p.NameVarEnvAPIKey, oldKey))
	}
}
