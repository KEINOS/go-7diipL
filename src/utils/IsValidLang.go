package utils

import "strings"

// IsValidLang は lang が QiiTrans の引数として有効な書式かチェックします.
func IsValidLang(lang string) bool {
	mapLang := GetListLang()

	result := mapLang[strings.ToLower(lang)]

	return result != ""
}
