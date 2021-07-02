package engine

import "os"

func (p *Properties) GetAPIKey() string {
	if p.apiKey != "" {
		return p.apiKey
	}

	return os.Getenv(p.NameVarEnvAPIKey)
}
