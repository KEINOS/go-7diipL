package app

// SetEngine メソッドは Engine フィールドに翻訳エンジンの新規インスタンスのポインタをセットします.
func (a *TApp) SetEngine(nameEngine string, cacheID ...string) (err error) {
	a.Engine, err = a.NewEngine(nameEngine, cacheID...)

	return err
}
