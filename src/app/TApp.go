package app

import "github.com/Qithub-BOT/QiiTrans/src/engines/engine"

// TApp はアプリの構造体です.
type TApp struct {
	Name    string             // アプリの公式名称（ヘルプの表示で使われます）
	Version string             // アプリのバージョン
	Engine  *engine.Properties // 翻訳に使うエンジン
	Argv    *TFlagOptions      // 現在保持しているフラグやオプションの値
}
