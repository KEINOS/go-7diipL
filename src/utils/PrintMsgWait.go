package utils

import (
	"fmt"
	"strings"
)

// PrintMsgWait は msg を改行なしで標準出力に出力するだけですが、そのメッセージ
// を打ち消すための関数を返します。
//
// 打ち消し関数を呼び出す前に別の標準出力があった場合は正常に打ち消されません。
func PrintMsgWait(msg string) (setoff func()) {
	lenMsg := len(msg)
	spacer := strings.Repeat(" ", lenMsg)

	fmt.Print(msg + "\x0D")

	return func() {
		fmt.Print("\x0D" + spacer + "\x0D")
	}
}
