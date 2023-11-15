package app

// TFlagOptions はアプリ（コマンド）のフラグやオプションの設定やその設定値およびヘルプの表示内容を定義した構造体です.
// アプリの bool オプション（フラグ）や string オプションなどは、ここで定義します.
//
//nolint:lll // allow long line due to help message
type TFlagOptions struct {
	APIKey         string `cli:"a,apikey"  dft:""                                                                                                                  usage:"翻訳に使うエンジンのアクセス・トークンを指定します"`
	CacheID        string `cli:"cache-id"  dft:""                                                                                                                  usage:"キャッシュの DB 名。異なる DB にキャッシュを保存したい場合に指定します"`
	NameEngine     string `cli:"e,engine"  dft:"deepl"                                                                                                             usage:"翻訳に使うエンジンを指定します"`
	UsageApp       string `cli:"-"` // アプリのヘルプ表示で使われる本文メッセージ
	ClearBeforeRun bool   `cli:"clear"     usage:"実行前にキャッシュを完全に削除します。（API の利用枠を消費します）"`
	Help           bool   `cli:"h,help"    usage:"ヘルプを表示します"`
	IsModeDebug    bool   `cli:"debug"     usage:"デバッグ情報を標準エラー出力に出力します"`
	IsNoCache      bool   `cli:"no-cache"  usage:"キャッシュを利用せずに翻訳 API から再取得します。（API の利用枠を消費します）"`
	IsPiped        bool   `cli:"-"` // パイプ渡しで値を受け取っているか
	IsVerbose      bool   `cli:"verbose"   usage:"中間翻訳も出力します。"`
	ShowInfo       bool   `cli:"info"      usage:"API のリクエスト可能な残数など、API 情報を出力します"`
	ShowInfoOnly   bool   `cli:"-"` // 他の引数がない場合は true になり API 情報のみの出力になります
	Version        bool   `cli:"v,version" usage:"アプリのバージョン情報を表示します"`
}
