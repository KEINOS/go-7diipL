package cache

// Get はキャッシュから key を探索し、その値を文字列で返します.
// キャッシュに key がない場合はエラーを返します.
func (c *TCache) Get(key string) (string, error) {
	if err := c.OpenDB(); err != nil {
		return "", err
	}

	defer c.CloseDB()

	// key をバイトデータにハッシュ化してキャッシュからキー探索
	value, err := c.GetValueInByte(c.hash64(key))

	return string(value), err
}
