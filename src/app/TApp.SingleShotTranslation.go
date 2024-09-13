package app

import (
	"fmt"

	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/gookit/color"
	"github.com/pkg/errors"
)

// SingleShotTranslation は標準入力から受け取ったテキストを翻訳する、単発翻訳用
// のメソッドです。
func (a *TApp) SingleShotTranslation(orderLang []string) (string, error) {
	// 標準入力から翻訳元のデータを取得
	input, err := utils.GetSTDIN()
	if err != nil {
		return "", errors.Wrap(err, "failed to get STDIN during single shot translation")
	}

	// 翻訳の実行
	listTranslated, err := a.Translate(orderLang, input)
	if err != nil {
		return "", errors.Wrap(err, "failed to translate during single shot translation")
	}

	lenTranslated := len(listTranslated)
	result := ""
	blue := color.Blue.Sprintf

	// Regular output
	if !a.Argv.IsVerbose {
		translated := listTranslated[lenTranslated-1]

		result = fmt.Sprint(blue(a.Prefix), " ", translated.Translated)

		return result, nil
	}

	// Verbose output
	for i := range lenTranslated {
		translated := listTranslated[i]
		prefix := a.Prefix

		// Print EN -> JA type prefix
		if i != lenTranslated-1 {
			prefix = fmt.Sprintf("%s -> %s:", translated.LangFrom, translated.LangTo)
		}

		result += fmt.Sprint(blue(prefix), " ", translated.Translated)
	}

	return result, nil
}
