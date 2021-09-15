package app

import "github.com/Qithub-BOT/QiiTrans/src/utils"

func (a *TApp) SingleShotTranslation(orderLang []string) (string, error) {
	// 標準入力から翻訳元のデータを取得
	input, err := utils.GetSTDIN()
	if err != nil {
		return "", err
	}

	// 翻訳の実行
	return a.Translate(orderLang, input)
}
