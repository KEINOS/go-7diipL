package main

const (
	// VersionAppDefault はアプリのバイナリをビルドする際にバージョン指定がない場合のアプリのバージョンです.
	VersionAppDefault = "dev"

	// MsgHelp はヘルプの表示内容です.
	//
	// "%s" は表示する際にアプリ名に自動置換されます。編集する際には文頭のスペースに注意.
	MsgHelp = `%s
  %s コマンドは文書作成支援ツールです。
  標準入力のテキストを、引数の言語から言語へ自動翻訳した結果を返します。

Usage:
  %s [Options] LangFrom LangTo [LangTo ...]

Example:
  $ # 日本語 → 英語に翻訳
  $ echo '私は、賛成の反対に同意なのだ' | %s ja en
  I agree to disagree.

  $ # 日本語 → 英語 → 日本語に翻訳
  $ echo '私は、賛成の反対に同意なのだ' | %s ja en ja
  同意しないことに同意します。

  $ # 日本語 → 英語 → スペイン語 → 中国語 → 日本語の順に翻訳
  $ echo '同意しないことに同意します。' | %s ja en es zh ja
  同意しないことに同意します。

LangFrom:
  標準入力から受け取るテキストの言語を指定します。（翻訳元の言語）

LangTo:
  翻訳先の言語を指定します。
  LangTo は複数指定可能で、1 つ前の言語から翻訳します。

翻訳エンジンについて
  現在は以下の翻訳 API に対応しています。（利用するには、無料もしくは有料アカウントの登録とアクセス・トークンの発行が必要です）

  deepl ... DeepL API: https://www.deepl.com/docs-api/ （デフォルトの翻訳エンジン）

翻訳エンジンとアクセス・トークンの設定について:
  %s は、翻訳 API のアクセストークンを環境変数から読み取ります。

  QIITRANS_API_KEY ... デフォルトで使うアクセス・トークンです。
  DEEPL_API_KEY ... DeepL 専用のアクセス・トークンです。セットされている場合は、QIITRANS_API_KEY より優先されます。
`
)
