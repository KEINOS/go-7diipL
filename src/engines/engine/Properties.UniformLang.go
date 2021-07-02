package engine

import (
	"strings"

	"github.com/Qithub-BOT/QiiTrans/src/utils"
)

// UniformLang は lang を API が検知可能な書式に変換します.
// もし lang が未対応の言語だった場合は、空の値を返します.
func (p *Properties) UniformLang(lang string) string {
	if !utils.IsValidLang(lang) {
		return ""
	}

	listLang := utils.GetListLang()

	return listLang[strings.ToLower(lang)]
}
