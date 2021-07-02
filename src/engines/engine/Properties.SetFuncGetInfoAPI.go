package engine

// SetFuncGetInfoAPI メソッドは翻訳用の関数を割り当てます.
func (p *Properties) SetFuncGetInfoAPI(getInfoFunc func(properties *Properties) (AccountInfo, error)) {
	p.getInfoAPI = getInfoFunc
}
