package app

import (
	"strings"
)

// GetMsgHelpUsage はヘルプ（使い方）のテンプレートから nameApp と nameCmd を置換した結果を返します.
func GetMsgHelpUsage(nameApp string, nameCmd string) string {
	usage := templateUsage

	usage = strings.ReplaceAll(usage, "%%nameapp%%", nameApp)
	usage = strings.ReplaceAll(usage, "%%namecmd%%", nameCmd)

	return usage
}

// MsgHelpUsage はアプリのヘルプ表示用のテンプレートです.
//
// メッセージ中の以下のタグは自動置換されます
//
//   "%%nameapp%%" ... アプリの公式名称
//   "%%namecmd%%" ... 実行バイナリ名（パスと拡張子除く）
//
// 注意: ヘルプはスペース・インデントです。タブ・インデントでレイアウト崩れが起きないように注意してください.
const templateUsage string = `%%nameapp%%
  %%namecmd%% コマンドは文書作成支援ツールです。
  標準入力のテキストを、引数の言語から言語へ自動翻訳した結果を返します。

Usage:
  %%namecmd%% [Options] LangFrom LangTo [LangTo ...]

Example:
  $ # 日本語 → 英語に翻訳
  $ echo '私は、賛成の反対に同意なのだ' | %%namecmd%% ja en
  I agree to disagree.

  $ # 日本語 → 英語 → 日本語に翻訳
  $ echo '私は、賛成の反対に同意なのだ' | %%namecmd%% ja en ja
  同意しないことに同意します。

  $ # 日本語 → 英語 → スペイン語 → 中国語 → 日本語の順に翻訳
  $ echo '同意しないことに同意します。' | %%namecmd%% ja en es zh ja
  同意しないことに同意します。

LangFrom:
  標準入力から受け取るテキストの言語を指定します。（翻訳元の言語）

LangTo:
  翻訳先の言語を指定します。
  LangTo は複数指定可能で、1 つ前の言語から翻訳します。

翻訳エンジンについて
  --engine オプションで翻訳 API を指定することができます。現在は以下の翻訳 API に対応しています。

    deepl ... DeepL Pro API を使ったエンジン（デフォルトの翻訳エンジン）
              URL: https://www.deepl.com/pro?cta=menu-login-signup (Free プランでも利用可能)
              認証キー確認先: https://www.deepl.com/pro-account/plan
              API情報: https://www.deepl.com/docs-api/

  注意: 利用するには、無料もしくは有料アカウントの登録とアクセストークン／認証キーの発行が必要です。

翻訳エンジンとアクセストークンの設定について:
  --apikey オプションでアクセストークン／認証キーが指定されていない場合、以下の環境変数から読み取ります。

  QIITRANS_API_KEY ... デフォルトで使うアクセストークンです。
  DEEPL_API_KEY ... DeepL 専用のアクセストークンです。セットされている場合は、QIITRANS_API_KEY より優先されます。
`
