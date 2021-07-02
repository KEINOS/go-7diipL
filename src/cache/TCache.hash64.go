package cache

import "github.com/Qithub-BOT/QiiTrans/src/utils"

// hash64 は content の内容を 64 バイトのハッシュ値に変換します.
func (c *TCache) hash64(content string) []byte {
	_, valByte, err := utils.Hash("blake3_512", content)
	utils.ExitOnErr(err)

	return valByte
}
