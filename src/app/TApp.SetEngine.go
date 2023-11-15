package app

import "github.com/pkg/errors"

// SetEngine メソッドは Engine フィールドに翻訳エンジンの新規インスタンスのポインタをセットします.
func (a *TApp) SetEngine(nameEngine string) error {
	ngin, err := a.NewEngine(nameEngine, a.cacheID...)
	if err == nil {
		a.Engine = ngin
	}

	return errors.Wrap(err, "failed to set translation engine")
}
