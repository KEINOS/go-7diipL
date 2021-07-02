//nolint: dupl // The expected values differ from other tests
package utils_test

import (
	"path/filepath"
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/stretchr/testify/assert"
)

func TestIsDir(t *testing.T) {
	pathDirTest := t.TempDir()

	genDummyDirAndFile(t, pathDirTest) // Defined in PathExists_test.go

	// Data Provider
	testCases := map[string]bool{
		"dir":                          true,
		"link-to-dir":                  true,
		"file.txt":                     false,
		".dotfile":                     false,
		"link-to-file.txt":             false,
		"unknown-dir":                  false,
		"unknown-file":                 false,
		".unknown-dotfile":             false,
		"unknown-link-to-file.txt":     false,
		"dir/unknown-dir":              false,
		"file.txt/unknown-dir":         false,
		".dotfile/unknown-dir":         false,
		"link-to-file.txt/unknown-dir": false,
		"link-to-dir/unknown-dir":      false,
	}

	for nameFileTest, expect := range testCases {
		pathFileTest := filepath.Join(pathDirTest, nameFileTest)
		actual := utils.IsDir(pathFileTest)

		assert.Equal(t, expect, actual, "Failed to detect path: %s", pathFileTest)
	}
}
