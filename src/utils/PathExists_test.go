package utils_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/stretchr/testify/assert"
)

func genDummyDirAndFile(t *testing.T, pathDirTemp string) {
	t.Helper()

	_ = os.RemoveAll(pathDirTemp)

	// create "dir"
	nameDir := "dir"
	pathDir := filepath.Join(pathDirTemp, nameDir)

	if err := os.MkdirAll(pathDir, 0o777); err != nil {
		t.Fatalf("Failed to create dir.\nMsg Error: %v", err)
	}

	// create "link-to-dir"
	nameDirSymlink := "link-to-dir"
	pathDirSymlink := filepath.Join(pathDirTemp, nameDirSymlink)
	pathDirTarget := pathDir

	if err := os.Symlink(pathDirTarget, pathDirSymlink); err != nil {
		t.Fatalf("Failed to create symbolic link of a directory.\nMsg Error: %v", err)
	}

	// create "file.txt"
	nameFile := "file.txt"
	pathFile := filepath.Join(pathDirTemp, nameFile)

	if _, err := os.Create(pathFile); err != nil {
		t.Fatalf("Failed to create file.\nMsg Error: %v", err)
	}

	// create ".dotfile"
	nameFileDot := ".dotfile"

	if _, err := os.Create(filepath.Join(pathDirTemp, nameFileDot)); err != nil {
		t.Fatalf("Failed to create dot file.\nMsg Error: %v", err)
	}

	// create "link-to-file.txt"
	nameFileSymlink := "link-to-file.txt"
	pathFileSymlink := filepath.Join(pathDirTemp, nameFileSymlink)
	pathFileTarget := pathFile

	if err := os.Symlink(pathFileTarget, pathFileSymlink); err != nil {
		t.Fatalf("Failed to create file.\nMsg Error: %v", err)
	}
}

func TestPathExists(t *testing.T) {
	pathDirTest := t.TempDir()

	genDummyDirAndFile(t, pathDirTest)

	// Data Provider
	testCases := map[string]bool{
		"dir":                          true,
		"link-to-dir":                  true,
		"file.txt":                     true,
		".dotfile":                     true,
		"link-to-file.txt":             true,
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
		actual := utils.PathExists(pathFileTest)

		assert.Equal(t, expect, actual, "Failed to detect path: %s", pathFileTest)
	}
}
