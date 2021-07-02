package engine

// GetQuotaLeft は API のリクエスト残量（翻訳可能文字数）を返します。有料アカウントや制限がない場合は -1 を返します.
func (p *Properties) GetQuotaLeft() (int, error) {
	info, err := p.getInfoAPI(p)

	return info.CharacterLeft, err
}
