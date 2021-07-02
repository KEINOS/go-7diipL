package main

import (
	"fmt"

	"github.com/Qithub-BOT/QiiTrans/src/utils"
)

// GetVersionApp は、アプリのバージョン情報を返します.
//
// デフォルトで dev ですが、アプリのバイナリをビルドする際に main.versionApp の値を指定する必要があります.
func GetVersionApp() string {
	return fmt.Sprintf("%s %s", utils.GetNameExe(), versionApp)
}
