package utils_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/stretchr/testify/assert"
)

const (
	pathDir            = "dir"
	pathDirSymlink     = "link-to-dir"
	pathFile           = "file.txt"
	pathFileDot        = ".dotfile"
	pathFileSymlink    = "link-to-file.txt"
	pathUnknownDir     = "unknown-dir"
	pathUnknownFile    = "unknown-file"
	pathUnknownDot     = ".unknown-dotfile"
	pathUnknownSymlink = "unknown-link-to-file.txt"
	pathDirSub         = "dir/unknown-dir"
	pathFileSub        = "file.txt/unknown-dir"
	pathDotSub         = ".dotfile/unknown-dir"
	pathSymlinkSub     = "link-to-file.txt/unknown-dir"
	pathDirSymlinkSub  = "link-to-dir/unknown-dir"
)

func genDummyDirAndFile(t *testing.T, pathDirTemp string) {
	t.Helper()

	_ = os.RemoveAll(pathDirTemp)

	// create "dir"
	pathDirAbs := filepath.Join(pathDirTemp, pathDir)

	err := os.MkdirAll(filepath.Clean(pathDirAbs), 0o750)
	if err != nil {
		t.Fatalf("Failed to create dir.\nMsg Error: %v", err)
	}

	// create "link-to-dir"
	pathDirSymlinkAbs := filepath.Join(pathDirTemp, pathDirSymlink)

	err = os.Symlink(pathDirAbs, pathDirSymlinkAbs)
	if err != nil {
		t.Fatalf("Failed to create symbolic link of a directory.\nMsg Error: %v", err)
	}

	// create "file.txt"
	pathFileAbs := filepath.Join(pathDirTemp, pathFile)

	_, err = os.Create(filepath.Clean(pathFileAbs))
	if err != nil {
		t.Fatalf("Failed to create file.\nMsg Error: %v", err)
	}

	// create ".dotfile"
	pathFileDotAbs := filepath.Join(pathDirTemp, pathFileDot)

	_, err = os.Create(filepath.Clean(pathFileDotAbs))
	if err != nil {
		t.Fatalf("Failed to create dot file.\nMsg Error: %v", err)
	}

	// create "link-to-file.txt"
	pathFileSymlinkAbs := filepath.Join(pathDirTemp, pathFileSymlink)

	err = os.Symlink(pathFileAbs, pathFileSymlinkAbs)
	if err != nil {
		t.Fatalf("Failed to create symbolic link of a file.\nMsg Error: %v", err)
	}
}

func TestPathExists(t *testing.T) {
	t.Parallel()

	pathDirTest := t.TempDir()

	genDummyDirAndFile(t, pathDirTest)

	// Data Provider
	testCases := map[string]bool{
		pathDir:            true,
		pathDirSymlink:     true,
		pathFile:           true,
		pathFileDot:        true,
		pathFileSymlink:    true,
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
		actual := utils.PathExists(pathFileTest)

		assert.Equal(t, expect, actual, "Failed to detect path: %s", pathFileTest)
	}
}
