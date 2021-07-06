package helperfunc_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/helperfunc"
	"github.com/stretchr/testify/assert"
)

func TestMockSTDIN(t *testing.T) {
	userInput := "foo bar"

	// stdin/パイプ からの入力のモックとリカバリ用関数取得
	funcDeferSTDIN := helperfunc.MockSTDIN(t, userInput)
	defer funcDeferSTDIN() // モックのリカバリ

	value, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		t.Fatalf("failed to read stdin during test")
	}

	expect := userInput
	actual := string(value)

	assert.Equal(t, expect, actual)
}
