package deepleng

import (
	"context"
	"os"

	"github.com/DaikiYamakawa/deepl-go"
	"github.com/Qithub-BOT/QiiTrans/src/engines/engine"
	"golang.org/x/xerrors"
)

// GetInfoAPI はアクセス・トークンのアカウント情報のうち、engine.AccountInfo に必要なものだけセットして返します.
func GetInfoAPI(p *engine.Properties) (engine.AccountInfo, error) {
	info := engine.AccountInfo{}

	if os.Getenv(p.NameVarEnvAPIKey) == "" {
		return info, xerrors.New("API key for DeepL not set")
	}

	if client, err := deepl.New(GetURLBaseAPI(p), nil); err == nil {
		accountStatus, err := client.GetAccountStatus(context.Background())

		if err == nil {
			// 利用可能の残文字数を取得
			limit := accountStatus.CharacterLimit
			used := accountStatus.CharacterCount

			// API からの取得値を設定
			info.CharacterLeft = limit - used
		}
	}

	return info, nil
}
