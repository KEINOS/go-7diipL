package app

// TTranslation は翻訳結果を保持するオブジェクトを定義します。
type TTranslation struct {
	LangFrom   string
	LangTo     string
	Original   string
	Translated string
}
