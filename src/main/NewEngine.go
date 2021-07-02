package main

import (
	"github.com/Qithub-BOT/QiiTrans/src/engines/deepleng"
	"github.com/Qithub-BOT/QiiTrans/src/engines/engine"
	"golang.org/x/xerrors"
)

// NewEngine は nameEngine で指定された翻訳エンジンのオブジェクト・ポインタを返します.
func NewEngine(nameEngine string) (*engine.Properties, error) {
	switch nameEngine {
	case "deepl":
		return deepleng.New(), nil
	}

	return nil, xerrors.New("未定義のエンジン名です: " + nameEngine)
}
