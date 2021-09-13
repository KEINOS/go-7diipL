package app

import (
	"strings"
)

// GetMsgHelpUsage はヘルプ（使い方）のテンプレートの定数 templateUsage から、name_app と name_cmd を置換した結果を返します.
func GetMsgHelpUsage(nameApp string, nameCmd string) string {
	if nameApp == "" {
		nameApp = NameDefault
	}

	usage := templateUsage

	usage = strings.ReplaceAll(usage, "%%name_app%%", nameApp)
	usage = strings.ReplaceAll(usage, "%%name_cmd%%", nameCmd)

	return usage
}
