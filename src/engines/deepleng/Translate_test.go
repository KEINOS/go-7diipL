package deepleng_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/engines/deepleng"
	"github.com/Qithub-BOT/QiiTrans/src/engines/engine"
	"github.com/stretchr/testify/assert"
)

func TestTranslate_no_env_set(t *testing.T) {
	e := engine.New(t.Name())
	defer e.Cache.ClearAll()

	e.NameVarEnvAPIKey = "UNEXISTING_DUMMY_KEY"

	_, err := deepleng.Translate(e, "input", "EN", "JA")

	assert.Error(t, err)
}
