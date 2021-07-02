package utils

// IsModeDebug はデバッグ・モードの場合に true を返します.
//
// Verbose 出力と異なり、ユーザ・サポートに必要な情報を出力する場合などに使います.
// デバッグ・モードの設定は SetModeDebut() を使います.
func IsModeDebug() bool {
	return isModeDebug
}
