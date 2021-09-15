package utils_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/helperfunc"
	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
)

func TestInteractSTDIN_not_terminal(t *testing.T) {
	stopWord := "stop"

	// utils.InteractSTDIN 内で利用するユーザ関数
	funcUser := func(input string) error {
		return nil
	}

	err := utils.InteractSTDIN(funcUser, stopWord)
	assert.Error(t, err)
}

func TestInteractSTDIN(t *testing.T) {
	stopWord := "stop"
	userInput := "foo bar\nhoge fuga\n" + stopWord + "\n"

	// utils.InteractSTDIN 内で利用するユーザ関数
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
		err := utils.InteractSTDIN(funcUser, stopWord)
		assert.NoError(t, err)
	})

	contain := "out:foo bar\nout:hoge fuga"

	assert.Contains(t, out, contain)
}

func TestInteractSTDIN_user_func_error(t *testing.T) {
	stopWord := "stop"
	expectError := "dummy error"
	userInput := "foo bar\nhoge fuga\n" + stopWord + "\n"

	// utils.InteractSTDIN 内で利用するユーザ関数
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
		err = utils.InteractSTDIN(funcUser, stopWord)
	})

	assert.Contains(t, out, stopWord)
	assert.Contains(t, out, "out:foo bar")
	assert.NotContains(t, out, "out:hoge fuga")

	assert.Error(t, err)
	assert.Contains(t, fmt.Sprintf("%v", err), expectError)
}
