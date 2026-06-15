package utils_test

import (
	"path/filepath"
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/stretchr/testify/assert"
)

func TestIsDir(t *testing.T) {
	t.Parallel()

	pathDirTest := t.TempDir()

	genDummyDirAndFile(t, pathDirTest) // Defined in PathExists_test.go

	// Data Provider
	testCases := map[string]bool{
		pathDir:            true,
		pathDirSymlink:     true,
		pathFile:           false,
		pathFileDot:        false,
		pathFileSymlink:    false,
		pathUnknownDir:     false,
		pathUnknownFile:    false,
		pathUnknownDot:     false,
		pathUnknownSymlink: false,
		pathDirSub:         false,
		pathFileSub:        false,
		pathDotSub:         false,
		pathSymlinkSub:     false,
		pathDirSymlinkSub:  false,
	}

	for nameFileTest, expect := range testCases {
		pathFileTest := filepath.Join(pathDirTest, nameFileTest)
		actual := utils.IsDir(pathFileTest)

		assert.Equal(t, expect, actual, "Failed to detect path: %s", pathFileTest)
	}
}
