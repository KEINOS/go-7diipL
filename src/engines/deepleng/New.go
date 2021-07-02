package deepleng

import "github.com/Qithub-BOT/QiiTrans/src/engines/engine"

const NameVarEnvAPIKey = "DEEPL_API_KEY"

// New は翻訳エンジン（DeepL）の新規インスタンスのポインタを返します.
func New(cacheID ...string) *engine.Properties {
	tmpEngine := engine.New(cacheID...)

	tmpEngine.SetDefault()                        // オブジェクトの初期化
	tmpEngine.SetFuncTrans(Translate)             // 翻訳関数の割り当て
	tmpEngine.SetFuncGetInfoAPI(GetInfoAPI)       // API アカウント情報取得関数の割り当て
	tmpEngine.NameVarEnvAPIKey = NameVarEnvAPIKey // DeepL 専用のアクセストークン取得・設定のための環境変数名

	return tmpEngine
}
