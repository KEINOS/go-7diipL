package utils

import "github.com/dustin/go-humanize"

// DelimitComma は int の値を 3 桁ごとにカンマを入れます（1000 -> 1,000）.
func DelimitComma(i int) string {
	return humanize.Comma(int64(i))
}
