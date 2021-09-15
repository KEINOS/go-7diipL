package app

// New はアプリの新規オプジェクトのポインタを返します.
//
// コマンド・オプションの --cache-id でも指定できるため、通常 cacheID は指定する必要はありません.
// テスト中、キャッシュがテスト間でバッティングしないようにキャッシュ ID を指定したい場合に利用します.
func New(cacheID ...string) *TApp {
	appTmp := new(TApp)

	appTmp.Argv = new(TFlagOptions)
	appTmp.cacheID = cacheID
	appTmp.Name = NameDefault
	appTmp.Prefix = PrefixDefault
	appTmp.StopWord = StopWordDefault
	appTmp.Version = VersionDefault

	// テスト用のフラグ. テスト時に強制的に戻り値を変える場合に利用します.
	appTmp.Force = map[string]bool{
		"FailPreRun": false, // True にすると PreRun が error を返します
		"IsNotPiped": false, // True にすると PreRun 内の IsPiped が false になります
		"NoTrans":    false, // True にすると TApp.Translate は翻訳せずに引数と同じ値を返します
		"TransError": false, // True にすると TApp.Translate はエラーを返しますï
	}

	return appTmp
}
