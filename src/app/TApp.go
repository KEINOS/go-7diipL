package app

import "github.com/Qithub-BOT/QiiTrans/src/engines/engine"

// TApp はアプリの構造体です.
type TApp struct {
	Argv    *TFlagOptions      // 現在保持しているフラグやオプションの値
	Engine  *engine.Properties // 翻訳に使うエンジン
	Name    string             // アプリの公式名称（ヘルプの表示で使われます）
	Version string             // アプリのバージョン
	cacheID []string           // キャッシュの DB 名
}
