package app_test

import (
	"runtime/debug"
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/app"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetVersion(t *testing.T) {
	appTest := app.New("", t.Name())

	{
		expect := "QiiTrans dev version"
		actual := appTest.GetVersion()

		assert.Equal(t, expect, actual, "it should be dev version by default")
	}
	{
		appTest.Version = ""

		expect := "QiiTrans dev version"
		actual := appTest.GetVersion()

		assert.Equal(t, expect, actual, "missing version should be as dev")
	}
	{
		appTest.Version = "1.0.0\n"

		expect := "QiiTrans v1.0.0"
		actual := appTest.GetVersion()

		assert.Equal(t, expect, actual, "the 'v' should be added and no line break")
	}
}

func TestGetVersion_via_build_info(t *testing.T) {
	// Backup and defer restore
	oldDebugReadBuildInfo := app.DebugReadBuildInfo
	defer func() {
		app.DebugReadBuildInfo = oldDebugReadBuildInfo
	}()

	// Expect version
	dummyVersion := "0.0.0-" + t.Name()

	// Mock
	//
	//nolint:nonamedreturns // allow named return dur to its readability
	app.DebugReadBuildInfo = func() (info *debug.BuildInfo, ok bool) {
		info = new(debug.BuildInfo)

		info.Path = t.Name()
		info.Main = debug.Module{
			Path:    "",
			Version: dummyVersion,
			Sum:     "",
			Replace: nil,
		}

		require.Equal(t, dummyVersion, info.Main.Version)

		return info, true
	}

	appTest := app.New("", t.Name())

	expect := "QiiTrans v" + dummyVersion
	actual := appTest.GetVersion()

	assert.Equal(t, expect, actual, "it should be the version via debug.ReadBuildInfo")
}
