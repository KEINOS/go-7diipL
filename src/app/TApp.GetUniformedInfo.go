package app

import (
	"fmt"

	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/pkg/errors"
)

// ForceErrorGetUniformedInfo が true の場合、GetUniformedInfo はエラーを返します.
//
// テスト時に強制的にエラーを返す必要がある場合のみ利用します.
var ForceErrorGetUniformedInfo = false

// GetUniformedInfo は API 情報を読みやすいように整えた状態で返します.
func (a *TApp) GetUniformedInfo() (string, error) {
	if ForceErrorGetUniformedInfo {
		return "", errors.New("forced to return error")
	}

	result := ""

	// 利用可能文字の残数取得
	quotaLeft, err := a.Engine.GetQuotaLeft()
	if err != nil {
		return "", errors.Wrap(err, "failed to get quota left")
	}

	result += fmt.Sprintf("[INFO]: 残り文字数: %v", utils.DelimitComma(quotaLeft))

	return result, nil
}
