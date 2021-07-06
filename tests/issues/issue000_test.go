package issues_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/stretchr/testify/assert"
)

// Sample of issue fix.
func TestIssue000(t *testing.T) {
	input := "foo bar buz"

	expect := []string{"foo bar buz"}
	actual := utils.SliceSentences(input)

	assert.Equal(t, expect, actual, "this sample test for issue fix should not fail")
}
