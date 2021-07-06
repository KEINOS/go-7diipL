package app

// New はアプリの新規オプジェクトのポインタを返します.
//
// コマンド・オプションの --cache-id でも指定できるため、通常 cacheID は指定する必要はありません.
// テスト中、キャッシュがテスト間でバッティングしないようにキャッシュ ID を指定したい場合に利用します.
func New(cacheID ...string) *TApp {
	appTmp := new(TApp)

	appTmp.Name = NameDefault
	appTmp.Version = VersionDefault
	appTmp.Argv = new(TFlagOptions)
	appTmp.cacheID = cacheID

	return appTmp
}
