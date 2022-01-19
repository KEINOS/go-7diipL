package app_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/app"
	"github.com/stretchr/testify/assert"
)

func TestNewTranslation(t *testing.T) {
	obj1 := app.NewTranslation("en", "ja", "sample1")
	obj2 := app.NewTranslation("en", "ja", "sample2")

	assert.NotSame(t, obj1, obj2, "pointers should notreference the same object.")
	assert.Equal(t, "en", obj1.LangFrom, "field LangFrom should have the 1st arg value on generation")
	assert.Equal(t, "ja", obj1.LangTo, "field LangTo should have the 2nd arg value on generation")
	assert.Equal(t, "sample1", obj1.Original, "field Original should have the 3rd arg value on generation")
	assert.Empty(t, obj1.Translated, "newly generated object should not have a value in Translated field")
}
