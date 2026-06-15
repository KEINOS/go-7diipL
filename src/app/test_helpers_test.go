package app_test

import (
	"os"
	"testing"
)

const (
	// flagInfo is the CLI flag used to display API quota information.
	flagInfo = "--info"
)

func requireDeepLAPIKey(t *testing.T) {
	t.Helper()

	if os.Getenv("DEEPL_API_KEY") == "" {
		t.Skip("DEEPL_API_KEY is not set")
	}
}
