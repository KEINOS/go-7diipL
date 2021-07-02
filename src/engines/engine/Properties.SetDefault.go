package engine

// SetDefault はプロパティ（オブジェクトのフィールド）値を初期値に設定します.
func (p *Properties) SetDefault() {
	p.TimeInterval = 1 // 1 リクエストごとに 1 秒待機させます
	p.Update = false   // API の翻訳結果を更新させずキャッシュさせます
}
