package utils

var isModeDebug bool = false

// SetModeDebug はデバッグ・モードのオン・オフを行います.
//
// 引数を true でセットすると GetModeDebug() は true を返します.
func SetModeDebug(flag bool) {
	isModeDebug = flag
}
