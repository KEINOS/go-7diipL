package main

import (
	"fmt"
	"strings"

	"github.com/Qithub-BOT/QiiTrans/src/utils"
)

// AutoHelp は cli.AutoHelper インターフェースの実装です.
//
// ヘルプを表示する際に使い方などの付加情報を設定しています.
func (argv *TFlag) AutoHelp() bool {
	if argv.Help {
		nameExe := utils.GetNameExe()

		header := MsgHelp
		header = strings.ReplaceAll(header, "%s", nameExe)

		fmt.Println(header)
	}

	return argv.Help
}
