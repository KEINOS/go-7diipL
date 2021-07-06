package app

import (
	"github.com/Qithub-BOT/QiiTrans/src/engines/deepleng"
	"github.com/Qithub-BOT/QiiTrans/src/engines/engine"
	"golang.org/x/xerrors"
)

// NewEngine メソッドは nameEngine で指定された翻訳エンジンの新規オブジェクト・ポインタを返します.
func (a *TApp) NewEngine(nameEngine string, cacheID ...string) (*engine.Properties, error) {
	switch nameEngine {
	case "deepl":
		return deepleng.New(cacheID...), nil
	}

	return nil, xerrors.New("未定義のエンジン名です: " + nameEngine)
}
