package app_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/app"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	appTest1 := app.New()
	appTest2 := app.New()

	assert.IsType(t, appTest1, appTest2, "it should create the same type")
	assert.NotSame(t, appTest1, appTest2, "pointers should not reference the same object")
}

func TestNew_default(t *testing.T) {
	appTest := app.New()

	// Check default name
	assert.Equal(t, app.NameDefault, appTest.Name, "it should be the default app name")

	// Check default version
	assert.Equal(t, app.VersionDefault, appTest.Version, "it should be the default app version")
}
