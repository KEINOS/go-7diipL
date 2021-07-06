package helperfunc_test

import (
	"os"
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/helperfunc"
	"github.com/stretchr/testify/assert"
)

func TestMockArgs(t *testing.T) {
	argsDummy := []string{
		"foo",
		"baf",
		"hoge",
		"fuga",
	}

	funcDefer := helperfunc.MockArgs(t, argsDummy)
	defer funcDefer()

	expect := argsDummy
	actual := os.Args[1:] // Args[0] is the command name

	assert.Equal(t, expect, actual)
}
