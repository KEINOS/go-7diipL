package utils_test

import (
	"fmt"
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/helperfunc"
	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/kami-zh/go-capturer"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInteractSTDIN_not_terminal(t *testing.T) {
	const (
		stopWord = "stop"
		prompt   = ">>>"
	)

	// utils.InteractSTDIN 内で利用するユーザ関数
	funcUser := func(_ string) error {
		return nil
	}

	err := utils.InteractSTDIN(funcUser, stopWord, prompt)
	require.Error(t, err)
}

func TestInteractSTDIN(t *testing.T) {
	const (
		stopWord = "stop"
		prompt   = ">>>"
	)

	userInput := "foo bar\nhoge fuga\n" + stopWord + "\n"

	// utils.InteractSTDIN 内で利用するユーザ関数
	//nolint:forbidigo // disable due to the nature of this function
	funcUser := func(input string) error {
		fmt.Println("out:" + input)

		return nil
	}

	// 標準入力をモック
	funcDefer := helperfunc.MockSTDIN(t, userInput)
	defer funcDefer()

	// utils.IsTerminal のモックとリカバリー
	utils.IsTerminalDummy = true
	defer func() { utils.IsTerminalDummy = false }()

	// テスト実行
	out := capturer.CaptureOutput(func() {
		err := utils.InteractSTDIN(funcUser, stopWord, prompt)
		require.NoError(t, err)
	})

	assert.Contains(t, out, prompt+"out:foo bar")
	assert.Contains(t, out, prompt+"out:hoge fuga")
}

func TestInteractSTDIN_user_func_error(t *testing.T) {
	const (
		stopWord = "stop"
		prompt   = ">>>"
	)

	expectError := "dummy error"
	userInput := "foo bar\nhoge fuga\n" + stopWord + "\n"

	// utils.InteractSTDIN 内で利用するユーザ関数
	//nolint:forbidigo // disable due to the nature of this function
	funcUser := func(input string) error {
		if input == "hoge fuga" {
			return errors.New(expectError)
		}

		fmt.Println("out:" + input)

		return nil
	}

	// 標準入力のモック
	funcDefer := helperfunc.MockSTDIN(t, userInput)
	defer funcDefer()

	// utils.IsTerminal のモックとリカバリー
	utils.IsTerminalDummy = true
	defer func() { utils.IsTerminalDummy = false }()

	var err error

	// エラーのテスト
	out := capturer.CaptureOutput(func() {
		err = utils.InteractSTDIN(funcUser, stopWord, prompt)

		require.Error(t, err)
		assert.Contains(t, err.Error(), expectError)
	})

	assert.Contains(t, out, stopWord)
	assert.Contains(t, out, "out:foo bar")
	assert.NotContains(t, out, "out:hoge fuga")
}
