package utils

import "github.com/dustin/go-humanize"

func DelimitComma(i int) string {
	return humanize.Comma(int64(i))
}
