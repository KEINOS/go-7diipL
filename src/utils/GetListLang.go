package utils

import "strings"

// listLang は利用可能な言語のマップです.
// 自動生成されるため GetListLang() か IsValidLang() を利用してください.
var listLang map[string]string

// GetListLang は、翻訳に使える言語の一覧をマップで返します.
//
// "map[検索]: 指定フォーマット" の形で返ってくるので、バリデーションに使います.
// "japanese" でキーを検索すると "ja"、"JA" でキーを検索しても "ja" が取得できます.
func GetListLang() map[string]string {
	if len(listLang) != 0 {
		return listLang
	}

	elements := map[string]string{
		"BG": "Bulgarian",
		"CS": "Czech",
		"DA": "Danish",
		"DE": "German",
		"EL": "Greek",
		"EN": "English",
		"ES": "Spanish",
		"ET": "Estonian",
		"FI": "Finnish",
		"FR": "French",
		"HU": "Hungarian",
		"IT": "Italian",
		"JA": "Japanese",
		"LT": "Lithuanian",
		"LV": "Latvian",
		"NL": "Dutch",
		"PL": "Polish",
		"PT": "Portuguese",
		"RO": "Romanian",
		"RU": "Russian",
		"SK": "Slovak",
		"SL": "Slovenian",
		"SV": "Swedish",
		"ZH": "Chinese",
	}

	elementMap := make(map[string]string)

	// [ja:JA japanese:JA] 形式のマップ作成
	for key, val := range elements {
		elementMap[strings.ToLower(key)] = key
		elementMap[strings.ToLower(val)] = key
	}

	listLang = elementMap

	return listLang
}
