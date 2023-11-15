package app

import (
	"fmt"
)

// AutoHelp は cli.AutoHelper インターフェースの実装です.
//
// この関数は cli.Run() から呼び出され、AutoHelp() が true を返した場合は cli.Run() は argv のオプションのヘルプを表示します.
func (argv *TFlagOptions) AutoHelp() bool {
	// ヘルプのオプションフラグが true の場合、アプリの usage を表示
	if argv.Help {
		//nolint:forbidigo // allow fmt due to help message
		fmt.Println(argv.UsageApp)
	}

	// argv.Help == true の場合、オプションフラグのヘルプが表示される
	return argv.Help
}
