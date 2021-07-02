package utils

import (
	"os"

	"golang.org/x/term"
)

// IsTerminalDummy はテスト時に標準入力からの受け取りをモックするために使われる値です.
// 値を true に設定すると IsTerminal() は強制的に true を返します。false に設定した場合は、自動検知が働きます.
var IsTerminalDummy bool = false

// IsTerminal はアプリがターミナル（コマンドライン）で実行されているか返します.
//
// テストでは必ず false になるため、テストでこの挙動を変えたい場合は IsTerminalDummy の値を true にします.
func IsTerminal() bool {
	if IsTerminalDummy {
		return true
	}

	return term.IsTerminal(int(os.Stdin.Fd()))
}
