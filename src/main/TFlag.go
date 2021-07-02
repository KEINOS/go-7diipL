package main

// TFlag は、コマンド引数を定義する構造体です.
type TFlag struct {
	APIKey         string `cli:"apikey" usage:"翻訳に使うエンジンのアクセス・トークンを指定します" dft:""`
	NameEngine     string `cli:"engine" usage:"翻訳に使うエンジンを指定します" dft:"deepl"`
	ClearBeforeRun bool   `cli:"clear" usage:"実行前にキャッシュを完全に削除します。（API の利用枠を消費します）"`
	Help           bool   `cli:"h,help" usage:"ヘルプを表示します"`
	IsModeDebug    bool   `cli:"debug" usage:"デバッグ情報を標準エラー出力に出力します"`
	IsNoCache      bool   `cli:"no-cache" usage:"キャッシュを利用せずに翻訳 API から再取得します。（API の利用枠を消費します）"`
	ShowInfo       bool   `cli:"info" usage:"API のリクエスト可能な残数など、API 情報を最後に出力します"`
	Version        bool   `cli:"v,version" usage:"アプリのバージョン情報を表示します"`
}
