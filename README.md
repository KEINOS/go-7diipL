[![All Tests](https://github.com/Qithub-BOT/QiiTrans/actions/workflows/merge-tests.yaml/badge.svg)](https://github.com/Qithub-BOT/QiiTrans/actions/workflows/merge-tests.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/Qithub-BOT/QiiTrans)](https://goreportcard.com/report/github.com/Qithub-BOT/QiiTrans)
[![Go Reference](https://pkg.go.dev/badge/github.com/Qithub-BOT/QiiTrans.svg)](https://pkg.go.dev/github.com/Qithub-BOT/QiiTrans)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/Qithub-BOT/QiiTrans/blob/main/LICENSE)

# `QiiTrans`

`qiitrans`コマンドは、文書作成支援ツールです。

`qiitrans`の機能は、標準入力から受け取ったテキストを翻訳して返すだけです。しかし、使い方を工夫すれば、機械翻訳で読まれても理解しやすい文章を作ることができます。

```shellsession
$ # 日本語 -> 英語に翻訳
$ echo 'おはよう。今日はいい天気ですね。' | qiitrans ja en
Good morning. It's a beautiful day.

$ # 英語 -> 日本語に翻訳
$ echo 'Good morning. It\'s a beautiful day.' | qiitrans en ja
おはようございます。今日はとてもいい天気ですね。
```

```shellsession
$ # 日本語 -> 英語 -> 日本語
$ echo '私は、賛成の反対に同意なのだ' | qiitrans ja en ja
同意しないことに同意します。

$ # 日本語 -> 英語 -> スペイン語 -> 中国語 -> 日本語
$ echo '同意しないことに同意します。' | qiitrans ja es en es zh ja
同意しないことに同意します。

$ echo '同意しないことに同意します。' | qiitrans ja en
I agree to disagree.
```

## インストール

Windows10, macOS, Linux (Intel/AMD 32bit/64bit, ARM v5,6,7 アーキテクチャ)用のバイナリを用意しています。

リリース・ページからダウンロードしたバイナリをパスの通ったディレクトリに設置後、実行権限を設定してください。

- [Releases](https://github.com/Qithub-BOT/QiiTrans/releases) | QiiTrans | Qithub @ GitHub

## コマンドの構文

```shellsession
$ qiitrans --help
qiitrans
  qiitrans コマンドは文書作成支援ツールです。
  標準入力のテキストを、引数の言語から言語へ自動翻訳した結果を返します。

Usage:
  qiitrans [Options] LangFrom LangTo [LangTo ...]

Example:
  $ # 日本語 → 英語に翻訳
  $ echo '私は、賛成の反対に同意なのだ' | qiitrans ja en
  I agree to disagree.

  $ # 日本語 → 英語 → 日本語に翻訳
  $ echo '私は、賛成の反対に同意なのだ' | qiitrans ja en ja
  同意しないことに同意します。

  $ # 日本語 → 英語 → スペイン語 → 中国語 → 日本語の順に翻訳
  $ echo '同意しないことに同意します。' | qiitrans ja en es zh ja
  同意しないことに同意します。

LangFrom:
  標準入力から受け取るテキストの言語を指定します。（翻訳元の言語）

LangTo:
  翻訳先の言語を指定します。
  LangTo は複数指定可能で、1 つ前の言語から翻訳します。

翻訳エンジンについて
  現在は以下の翻訳 API に対応しています。
  利用するには、無料もしくは有料アカウントの登録とアクセス・トークンの発行が必要です。

  deepl ... DeepL API: https://www.deepl.com/docs-api/ （デフォルトの翻訳エンジン）

翻訳エンジンとアクセス・トークンの設定について:
  qiitrans は、翻訳 API のアクセストークンを環境変数から読み取ります。

  QIITRANS_API_KEY ... デフォルトで使うアクセス・トークンです。
  DEEPL_API_KEY ... DeepL 専用のアクセス・トークンです。セットされている場合は、QIITRANS_API_KEY より優先されます。


Options:

      --apikey           翻訳に使うエンジンのアクセス・トークンを指定します
      --engine[=deepl]   翻訳に使うエンジンを指定します
      --clear            実行前にキャッシュを完全に削除します。（API の利用枠を消費します）
  -h, --help             ヘルプを表示します
      --debug            デバッグ情報を標準エラー出力に出力します
      --no-cache         キャッシュを利用せずに翻訳 API から再取得します。（API の利用枠を消費します）
      --info             API のリクエスト可能な残数など、API 情報を最後に出力します
```

## アクセス・トークン

`QiiTrans` は、翻訳 API のアクセス・トークンを必要とします。以下のいずれかで指定し、最初に該当したアクセス・トークンを利用します。

1. コマンド・オプションの `--apikey` で指定されたトークン。
2. 環境変数に、使用している翻訳エンジン専用の変数がセットされている場合は、その値。
    - `DEEPL_API_KEY` ... `--engine=deepl` （デフォルトの翻訳エンジン）
3. `QIITRANS_API_KEY` がセットされている場合は、その値。

## その他の環境変数

`QiiTrans` は、以下の環境変数がセットされていた場合は、デフォルト値として利用します。

- `QIITRANS_ENGINE`: 翻訳に使うエンジンです。現在は "`deepl`" のみが指定可能です。
- `QIITRANS_API_KEY`: 翻訳に使うエンジンの APIキー（アクセストークン）です。

## コントリビュート

- [質問・要望・改善案・提案](https://github.com/Qithub-BOT/QiiTrans/discussions) @ Discussions
- [不具合報告（要再現テストあり）](https://github.com/Qithub-BOT/QiiTrans/issues) @ Issues
- [TODO](https://github.com/Qithub-BOT/QiiTrans/issues) @ Issues
  - [WIP](https://github.com/Qithub-BOT/QiiTrans/pulls) | Draft PR @ Pull Requests

## ライセンスと著作権

- License: [MIT](https://github.com/Qithub-BOT/QiiTrans/blob/master/LICENSE.md)
- Copyright: (c) [The QiiTrans Contributors](https://github.com/Qithub-BOT/QiiTrans/graphs/contributors)
