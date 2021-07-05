package app

import (
	"strings"
)

// Translate は、orderLang の順に inputText を翻訳した結果を返します.
func (a *TApp) Translate(orderLang []string, inputText string) (string, error) {
	var err error

	transText := ""
	langFrom := ""
	langTo := ""

	// orderLang 順に再帰的に翻訳
	for i, nameLang := range orderLang {
		nameLang = strings.ToUpper(nameLang) // API に合わせて言語を大文字に統一

		// 最初の言語をセット
		if i == 0 {
			langFrom = nameLang
			transText = inputText

			continue
		}

		langTo = nameLang

		// from -> to へ翻訳
		transText, _, err = a.Engine.Translate(transText, langFrom, langTo)
		if err != nil {
			return "", err
		}

		langFrom = langTo
	}

	return transText, err
}
