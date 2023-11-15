package engine

import (
	"strings"

	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/pkg/errors"
)

// Translate は翻訳エンジンから割り当てられた翻訳関数を実行し、inText を langFrom から langTo に翻訳した結果を返します.
// この値はキャッシュされます。キャッシュを更新したい場合は Update フィールド（プロパティ）を true にセットしてください.
//
//nolint:nonamedreturns // use named return for better readability
func (p *Properties) Translate(inTxt string, langFrom string, langTo string) (outText string, isCache bool, err error) {
	// fix issue #8
	if strings.TrimSpace(inTxt) == "" {
		return inTxt, false, nil
	}

	var (
		result string
		key    string
	)

	key = langFrom + langTo + inTxt
	isCache = false

	// キャッシュがある場合はキャッシュを返します。Update プロパティが true の場合は更新します
	result, err = p.Cache.Get(key)
	if err == nil || !p.Update {
		if strings.TrimSpace(result) != "" {
			utils.LogDebug("%s -> %s: キャッシュ: %s", langFrom, langTo, result)
			utils.LogDebug("Dir path: %s", p.Cache.PathDirTemp)

			isCache = true

			return result, isCache, nil
		}
	}

	// 翻訳エンジンの翻訳関数を叩いて、翻訳結果を取得します
	result, err = p.translate(p, inTxt, langFrom, langTo)
	if err != nil {
		return "", isCache, errors.Wrap(err, "failed to translate")
	}

	utils.LogDebug("%s -> %s: 新規取得: %s", langFrom, langTo, result)

	// 翻訳結果をキャッシュします
	err = p.Cache.Set(key, result)

	return result, isCache, errors.Wrap(err, "failed to set cache")
}
