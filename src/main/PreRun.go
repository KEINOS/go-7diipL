package main

import (
	"github.com/Qithub-BOT/QiiTrans/src/utils"
)

// PreRun はアプリの初期化を行います.
func PreRun(argv *TFlag) error {
	// デバッグ・モードのセット
	utils.SetModeDebug(argv.IsModeDebug)

	return nil
}
