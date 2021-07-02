package utils_test

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
)

func TestInteractSTDIN_not_terminal(t *testing.T) {
	stopWord := "stop"

	// utils.InteractSTDIN 内で利用するユーザ関数
	funcUser := func(input string) (string, error) {
		return "out:" + input, nil
	}

	err := utils.InteractSTDIN(funcUser, stopWord)
	assert.Error(t, err)
}

func TestInteractSTDIN(t *testing.T) {
	stopWord := "stop"

	input := "foo bar\nhoge fuga\n" + stopWord + "\n"

	// utils.InteractSTDIN 内で利用するユーザ関数
	funcUser := func(input string) (string, error) {
		return "out:" + input, nil
	}

	// 標準入力のダミー用ファイルポインタ作成
	tmpFile, funcDefer := mockSTDIN(t, input)

	// ダミーのファイル・ポインタ clean up
	defer funcDefer()

	// os.Stdin のバックアップとリカバリー
	oldStdin := os.Stdin

	defer func() { os.Stdin = oldStdin }()

	// os.Stdin をモック
	os.Stdin = tmpFile

	// utils.IsTerminal のモックとリカバリー
	utils.IsTerminalDummy = true

	defer func() { utils.IsTerminalDummy = false }()

	out := capturer.CaptureOutput(func() {
		err := utils.InteractSTDIN(funcUser, stopWord)
		assert.NoError(t, err)
	})

	contain := "out:foo bar\nout:hoge fuga"

	assert.Contains(t, out, contain)
}

func TestInteractSTDIN_user_func_error(t *testing.T) {
	stopWord := "stop"
	expectError := "dummy error"

	input := "foo bar\nhoge fuga\n" + stopWord + "\n"

	// utils.InteractSTDIN 内で利用するユーザ関数
	funcUser := func(input string) (string, error) {
		if input == "hoge fuga" {
			return "", errors.New(expectError)
		}

		return "out:" + input, nil
	}

	// 標準入力のダミー用ファイルポインタ作成
	tmpFile, funcDefer := mockSTDIN(t, input)

	// ダミーのファイル・ポインタ clean up
	defer funcDefer()

	// os.Stdin のバックアップとリカバリー
	oldStdin := os.Stdin

	defer func() { os.Stdin = oldStdin }()

	// os.Stdin をモック
	os.Stdin = tmpFile

	// utils.IsTerminal のモックとリカバリー
	utils.IsTerminalDummy = true

	defer func() { utils.IsTerminalDummy = false }()

	var err error

	out := capturer.CaptureOutput(func() {
		err = utils.InteractSTDIN(funcUser, stopWord)
	})

	assert.Contains(t, out, stopWord)
	assert.Contains(t, out, "out:foo bar")
	assert.NotContains(t, out, "out:hoge fuga")

	assert.Error(t, err)
	assert.Contains(t, fmt.Sprintf("%v", err), expectError)
}
