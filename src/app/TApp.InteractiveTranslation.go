package app

import (
	"fmt"

	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/gookit/color"
)

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
