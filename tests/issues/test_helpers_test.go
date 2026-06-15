package issues_test

import (
	"os"
	"testing"
)

func requireDeepLAPIKey(t *testing.T) {
	t.Helper()

	if os.Getenv("DEEPL_API_KEY") == "" {
		t.Skip("DEEPL_API_KEY is not set")
	}
}
