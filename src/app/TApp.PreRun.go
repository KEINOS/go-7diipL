package app

import (
	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"golang.org/x/xerrors"
)

// ForceFailPreRun は PreRun メソッドを強制的に失敗させる（1 を返す）ためのフラグです.
// 主にテスト目的で使われます.
var ForceFailPreRun = false

// PreRun は Run の本体処理を行う前にフラグ、オプションなどの引数のセットなどを行います.
//
// テスト目的で強制的に失敗させる場合は ForceFailPreRun を true に設定します.
func (a *TApp) PreRun() error {
	if ForceFailPreRun {
		return xerrors.New("PreRun was forced to fail")
	}

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
