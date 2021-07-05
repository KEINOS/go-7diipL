package app

import (
	"fmt"
	"strings"
)

// GetVersion メソッドはアプリ名を含めたバージョン情報を返します.
//
// Version フィールドの値が "v" で始まらない場合は、頭に付け加えます.
// Version フィールドの値が、空もしくは "dev" の場合は "dev version" になります.
func (a *TApp) GetVersion() string {
	nameApp := strings.TrimSpace(a.Name)
	verApp := strings.TrimSpace(a.Version)

	// dev version
	if verApp == "" || verApp == VersionDefault {
		return fmt.Sprintf("%s %s version", nameApp, VersionDefault)
	}

	// missing "v"
	if !strings.HasPrefix(verApp, "v") {
		verApp = "v" + verApp
	}

	return fmt.Sprintf("%s %s", nameApp, verApp)
}
