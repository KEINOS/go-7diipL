package deepleng

import "github.com/Qithub-BOT/QiiTrans/src/engines/engine"

// NameVarEnvAPIKey は環境変数の変数名で、DeepL の認証キー（アクセストークン）用の変数名です.
//
//nolint:gosec // It is not a hardcoded credentials
const NameVarEnvAPIKey = "DEEPL_API_KEY"

// New は翻訳エンジン（DeepL）の新規インスタンスのポインタを返します.
func New(cacheID ...string) *engine.Properties {
	tmpEngine := engine.New(cacheID...)

	tmpEngine.SetDefault() // オブジェクトの初期化

	tmpEngine.SetFuncTrans(Translate)       // 翻訳関数の割り当て
	tmpEngine.SetFuncGetInfoAPI(GetInfoAPI) // API アカウント情報取得関数の割り当て

	tmpEngine.NameEngine = "deepl"                // エンジン名（利用翻訳 API 名）
	tmpEngine.NameVarEnvAPIKey = NameVarEnvAPIKey // DeepL 専用のアクセストークン取得・設定のための環境変数名

	return tmpEngine
}
