package issues_test

import (
	"os"
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/engines/deepleng"
)

func requireDeepLAPIKey(t *testing.T) {
	t.Helper()

	if os.Getenv("DEEPL_API_KEY") == "" {
		t.Skip("DEEPL_API_KEY is not set")
	}

	eng := deepleng.New(t.Name())
	defer eng.Cache.ClearAll()

	_, _, err := eng.Translate("auth check", "EN", "JA")
	if err != nil {
		t.Skipf("DEEPL_API_KEY is not usable for integration tests: %v", err)
	}
}
