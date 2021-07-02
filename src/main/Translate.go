package main

import (
	"strings"

	"github.com/Qithub-BOT/QiiTrans/src/engines/engine"
)

// Translate は e の翻訳エンジンを使って orderLang で指定された言語の順に inputText を翻訳し、結果を返します.
func Translate(e *engine.Properties, orderLang []string, inputText string) (string, error) {
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
		transText, _, err = e.Translate(transText, langFrom, langTo)
		if err != nil {
			return "", err
		}

		langFrom = langTo
	}

	return transText, err
}
