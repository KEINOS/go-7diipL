package app

import "github.com/Qithub-BOT/QiiTrans/src/utils"

func (a *TApp) PreRun() error {
	// デバッグ・モードのセット
	utils.SetModeDebug(a.Argv.IsModeDebug)
	a.Argv.UsageApp = GetMsgHelpUsage(a.Name, utils.GetNameExe())

	return nil
}
