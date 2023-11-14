package engine

// SetFuncTrans メソッドは翻訳用の関数を割り当てます.
func (p *Properties) SetFuncTrans(transFunc func(
	properties *Properties,
	inputText string,
	langFrom string,
	langTo string,
) (string, error),
) {
	p.translate = transFunc
}
