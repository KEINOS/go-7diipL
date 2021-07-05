package app

import "github.com/Qithub-BOT/QiiTrans/src/utils"

// SetHelpMsg はヘルプのテンプレートに現在のアプリ名及び実行ファイル名を流し込みセットします.
func (argv *TFlagOptions) SetHelpMsg() {
	argv.UsageApp = GetMsgHelpUsage(argv.NameApp, utils.GetNameExe())
}
