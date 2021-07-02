package cache

import "path/filepath"

// GetPathDirCache は現在のキャッシュ先のディレクトリを返します.
func (c *TCache) GetPathDirCache() string {
	return filepath.Join(c.PathDirTemp, c.NameDirCache)
}
