package engine

import "os"

// GetAPIKey はコマンドのオプションで指定されたアクセス・トークンを返します.
// オプションで指定がされていない場合は、環境変数から読み取って返します.
func (p *Properties) GetAPIKey() string {
	if p.apiKey != "" {
		return p.apiKey
	}

	return os.Getenv(p.NameVarEnvAPIKey)
}
