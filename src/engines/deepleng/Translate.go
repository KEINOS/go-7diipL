package deepleng

import (
	"context"
	"os"

	"github.com/DaikiYamakawa/deepl-go"
	"github.com/Qithub-BOT/QiiTrans/src/engines/engine"
	"github.com/pkg/errors"
	"golang.org/x/xerrors"
)

// Translate は inText を langFrom から langTo に DeepL の翻訳 API を使って翻訳した結果を返します.
//
// この関数は結果をキャッシュしません。別途、呼び出し元でキャッシュを行ってください.
//
//nolint:nonamedreturns // use named return for better readability
func Translate(prop *engine.Properties, inText string, langFrom string, langTo string) (result string, err error) {
	if os.Getenv(prop.NameVarEnvAPIKey) == "" {
		return "", xerrors.New("API key for DeepL not set")
	}

	var (
		client            *deepl.Client
		translateResponse *deepl.TranslateResponse
	)

	client, err = deepl.New(GetURLBaseAPI(prop), nil)
	if err == nil {
		translateResponse, err = client.TranslateSentence(context.Background(), inText, langFrom, langTo)
		if err == nil {
			result = translateResponse.Translations[0].Text // 最もマッチする結果を返す
		}
	}

	return result, errors.Wrap(err, "failed to translate")
}
