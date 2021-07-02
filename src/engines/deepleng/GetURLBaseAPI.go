package deepleng

import "github.com/Qithub-BOT/QiiTrans/src/engines/engine"

// GetURLBaseAPI は DeepL の翻訳 API のエンドポイントのドメインを返します.
// 無料・有料アカウントでドメインは異なるため、アクセス・トークン（認証キー）から自動検知して返します.
func GetURLBaseAPI(p *engine.Properties) string {
	apiDomain := "https://api.deepl.com"

	if p.IsAccountFree || IsAPIKeyFree(p.GetAPIKey()) {
		apiDomain = "https://api-free.deepl.com"
	}

	return apiDomain
}
