package app

import "github.com/Qithub-BOT/QiiTrans/src/utils"

func (a *TApp) PreRun() error {
	// デバッグ・モードのセット
	utils.SetModeDebug(a.Argv.IsModeDebug)
	a.Argv.UsageApp = GetMsgHelpUsage(a.Name, utils.GetNameExe())

	// 翻訳エンジンのセット
	if err := a.SetEngine(a.Argv.NameEngine); err != nil {
		return err
	}

	// キャッシュの可否をセット
	a.Engine.Update = a.Argv.IsNoCache

	return nil
}
