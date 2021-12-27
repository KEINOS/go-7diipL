package app

import (
	"fmt"

	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/gookit/color"
)

// InteractiveTranslation は、対話式（標準入力がない場合）の連続的な翻訳を行う
// メソッドです。
func (a *TApp) InteractiveTranslation(orderLang []string) error {
	stopWord := a.StopWord
	translator := func(input string) error {
		output, err := a.Translate(orderLang, input)
		if err != nil {
			return err
		}

		prefix := color.Blue.Sprintf(a.Prefix)

		fmt.Println(prefix, output)

		return nil
	}

	return utils.InteractSTDIN(translator, stopWord)
}
