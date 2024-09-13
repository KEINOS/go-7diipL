package app

import (
	"fmt"

	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/gookit/color"
	"github.com/pkg/errors"
)

// InteractiveTranslation は、対話式（標準入力がない場合）の連続的な翻訳を行う
// メソッドです。
//
//nolint:forbidigo // allow fmt due to its function
func (a *TApp) InteractiveTranslation(orderLang []string) error {
	stopWord := a.StopWord
	prompt := a.Prompt
	translator := func(input string) error {
		listTranslated, err := a.Translate(orderLang, input)
		if err != nil {
			return errors.Wrap(err, "translator failed")
		}

		utils.LogDebug(fmt.Sprintf("%#v", listTranslated))

		lenTranslated := len(listTranslated)
		blue := color.Blue.Sprintf

		// Usual print
		if !a.Argv.IsVerbose {
			translated := listTranslated[lenTranslated-1]

			fmt.Println(blue(a.Prefix), translated.Translated)

			return nil
		}

		// Verbose print
		for i := range lenTranslated {
			translated := listTranslated[i]
			prefix := a.Prefix

			if i != lenTranslated-1 {
				prefix = fmt.Sprintf("%s -> %s:", translated.LangFrom, translated.LangTo)
			}

			fmt.Println(blue(prefix), translated.Translated)
		}

		return nil
	}

	err := utils.InteractSTDIN(translator, stopWord, prompt)

	return errors.Wrap(err, "failed during interactive translation")
}
