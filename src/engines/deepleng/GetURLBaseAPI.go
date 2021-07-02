package deepleng

import "github.com/Qithub-BOT/QiiTrans/src/engines/engine"

func GetURLBaseAPI(p *engine.Properties) string {
	apiDomain := "https://api.deepl.com"

	if p.IsAccountFree || IsAPIKeyFree(p.GetAPIKey()) {
		apiDomain = "https://api-free.deepl.com"
	}

	return apiDomain
}
