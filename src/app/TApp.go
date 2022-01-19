package app

import "github.com/Qithub-BOT/QiiTrans/src/engines/engine"

// TApp はアプリの構造体です.
type TApp struct {
	Argv     *TFlagOptions
	Engine   *engine.Properties
	Force    map[string]bool
	Name     string
	Prefix   string // 対話モード時に翻訳結果の前につける接頭辞
	Prompt   string // 対話モード時の待受プロンプト文字（デフォルト: >>>）
	StopWord string // 対話モード時に終了を伝える単語（デフォルト: q）
	Version  string
	cacheID  []string
}
