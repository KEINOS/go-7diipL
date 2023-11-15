package app

import (
	"github.com/Qithub-BOT/QiiTrans/src/engines/deepleng"
	"github.com/Qithub-BOT/QiiTrans/src/engines/engine"
	"github.com/pkg/errors"
)

// NewEngine メソッドは nameEngine で指定された翻訳エンジンの新規オブジェクト・ポインタを返します.
func (a *TApp) NewEngine(nameEngine string, cacheID ...string) (*engine.Properties, error) {
	//nolint:gocritic // allow switch-case due to future expansion
	switch nameEngine {
	case "deepl":
		return deepleng.New(cacheID...), nil
	}

	return nil, errors.New("未定義のエンジン名です: " + nameEngine)
}
