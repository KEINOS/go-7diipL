package app

import (
	"os"

	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/pkg/errors"
	"golang.org/x/xerrors"
)

// PreRun は Run の本体処理を行う前にフラグ、オプションなどの引数のセットなどを行います.
func (a *TApp) PreRun() error {
	if a.Force["FailPreRun"] {
		return xerrors.New("PreRun was forced to fail")
	}

	// デバッグ・モードのセット
	utils.SetModeDebug(a.Argv.IsModeDebug)

	a.Argv.UsageApp = GetMsgHelpUsage(a.Name, utils.GetNameExe())

	// 翻訳エンジンのセット
	if err := a.SetEngine(a.Argv.NameEngine); err != nil {
		return errors.Wrap(err, "failed to set translation engine")
	}

	// キャッシュの可否をセット
	a.Engine.Update = a.Argv.IsNoCache

	// パイプ渡しで値を受け取っているかをセット
	a.Argv.IsPiped = false

	stat, err := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		a.Argv.IsPiped = true
	}

	// テスト用の強制フラグ
	if a.Force["IsNotPiped"] {
		a.Argv.IsPiped = false
	}

	return errors.Wrap(err, "failed to set pre-run")
}
