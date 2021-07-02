package engine

// AccountInfo は、利用中の翻訳 API の必要最低限の情報を保持する構造体です.
type AccountInfo struct {
	CharacterLeft int // = 利用可能文字数 - 利用済み文字数
}
