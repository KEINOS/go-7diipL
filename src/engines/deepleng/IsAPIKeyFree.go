package deepleng

import "strings"

// IsAPIKeyFree は apikey が無料枠のアクセストークンか返します.
func IsAPIKeyFree(apikey string) bool {
	// DeepL の API の無料枠アクセストークンは末尾に ":fx" が付きます
	suffix := ":fx"

	return strings.HasSuffix(apikey, suffix)
}
