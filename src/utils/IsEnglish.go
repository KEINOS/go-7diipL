package utils

import (
	"strings"

	"github.com/abadojack/whatlanggo"
)

// IsEnglish は input が英文であると予測した場合に true を返します.
func IsEnglish(input string) bool {
	listItem := strings.Split(input, "\n")
	isEng := true

	for _, item := range listItem {
		langInfo := whatlanggo.Detect(item) // テキストの言語を検知

		if langInfo.Lang.Iso6393() != whatlanggo.Eng.Iso6393() {
			return false
		}
	}

	return isEng
}
