package app

import (
	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/mkideal/cli"
)

// ForceFailRun は Run メソッドを強制的に失敗させる（1 を返す）ためのフラグです.
// 主にテスト目的で使われます.
var ForceFailRun bool = false

// Run メソッドはアプリを実行します. 実行に成功した場合は nil を、失敗した場合は error を返します.
//
// テスト目的で強制的に失敗させる場合は ForceFail を true に設定します.
func (a *TApp) Run() int {
	if ForceFailRun {
		return utils.FAILURE
	}

	tmpFlag := new(TFlagOptions)
	tmpFlag.SetHelpMsg()

	// アプリ本体のメソッド TApp.CliRun を cli.Run に割り当て
	return cli.Run(tmpFlag, a.CliRun)
}
