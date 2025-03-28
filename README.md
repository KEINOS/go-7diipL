> __Note__: The repository is currently in the process of a major redesign and conversion to an English-language base.

---

[![All Tests](https://github.com/Qithub-BOT/QiiTrans/actions/workflows/merge-tests.yml/badge.svg)](https://github.com/Qithub-BOT/QiiTrans/actions/workflows/merge-tests.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/Qithub-BOT/QiiTrans)](https://goreportcard.com/report/github.com/Qithub-BOT/QiiTrans)
[![CodeQL](https://github.com/Qithub-BOT/QiiTrans/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/Qithub-BOT/QiiTrans/actions/workflows/codeql-analysis.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/Qithub-BOT/QiiTrans.svg)](https://pkg.go.dev/github.com/Qithub-BOT/QiiTrans)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/Qithub-BOT/QiiTrans/blob/main/LICENSE.md)
[![](https://shields.io/badge/GitHub-Codespaces%20%E5%AF%BE%E5%BF%9C-blue?logo=github&style=flat)](https://docs.github.com/en/codespaces "このリポジトリは GitHub Codespaces に対応しています")

# `QiiTrans`

`qiitrans` コマンドは文書作成の支援ツールです。

## Install

<details><summary>macOS/Linux/Windows (x86_64,Intel,AMD / ARM64 / Armv 5,6,7)</summary><br>

### Homebrew

Linux, macOS の場合は [Homebrew](https://brew.sh/) でインストールできます。（macOS は Intel と M1 にも対応しています）

```bash
brew install qithub-bot/apps/qiitrans
```

### 手動インストール

Windows10, macOS, Linux 用の単体バイナリを用意しています。

リリース・ページから OS と CPU に対応したバイナリをダウンロードし、パスの通ったディレクトリに実行権限付けて設置してください。

- 対応アーキテクチャ: Intel/AMD 32bit/64bit, ARM64, ARM v5,6,7
- [Releases](https://github.com/Qithub-BOT/QiiTrans/releases) | QiiTrans | Qithub @ GitHub

</details>

## Usage

`qiitrans` の機能は、受け取ったテキストを翻訳して返すだけです。

```shellsession
$ # Requirements: DeepL Access Token
$ env | grep DEEPL_API_KEY
DEEPL_API_KEY=12345678-ffff-ffff-ffff-123456789abc:fx
```
```shellsession
$ # 日本語 → English に翻訳
$ echo 'おはよう。今日はいい天気ですね。' | qiitrans ja en
再翻訳: Good morning. It's a beautiful day.

$ # English → 日本語に翻訳
$ echo 'Good morning. It\'s a beautiful day.' | qiitrans en ja
再翻訳: おはようございます。今日はとてもいい天気ですね。
```

しかし、使い方を工夫すれば、機械翻訳で読まれても理解しやすい文章を作ることができます。

```shellsession
$ # English → Español → 日本語 → 中文 → English
$ echo 'Difficult the document to read, it is.' | qiitrans en es ja zh en
再翻訳: The document is difficult to read.

$ echo 'The document is difficult to read.' | qiitrans en es ja zh en
再翻訳: The document is difficult to read.
```

```shellsession
$ # 日本語 → English → 日本語
$ echo '私は、賛成の反対に同意なのだ。' | qiitrans ja en ja
再翻訳: 同意しないことに同意します。

$ # 日本語 → Español → English → Español → 中文 → 日本語
$ echo '同意しないことに同意します。' | qiitrans ja es en es zh ja
再翻訳: 同意しないことに同意します。

$ echo '同意しないことに同意します。' | qiitrans ja en
再翻訳: I agree to disagree.
```

### REPL（対話モード）

標準入力がない場合は対話モードになります。ストップワードが入力されるまで翻訳を続けます。

```shellsession
$ qiitrans en es ja zh en
- ストップワード: q
>>> Difficult the code to read, it is.
再翻訳: This is a difficult code to read.
>>> q
q を検知しました。対話モードを終了します。
```

### Advanced Usage

`--verbose` オプションを付けると、途中の翻訳も見ることができます。

```shellsession
$ # One-liner
$ echo 'Difficult the document to read, it is.' | qiitrans en es ja zh en --verbose
EN -> ES: Difícil de leer es el documento.
ES -> JA: ドキュメントが読みにくい。
JA -> ZH: 文件难以阅读。
再翻訳: The document is difficult to read.
```
```shellsession
$ # Interactive Mode
$ qiitrans en es ja zh en --verbose
- ストップワード: q
>>> Difficult the document to read, it is.
EN -> ES: Difícil de leer es el documento.
ES -> JA: ドキュメントが読みにくい。
JA -> ZH: 文件难以阅读。
再翻訳: The document is difficult to read.
>>> q
q を検知しました。対話モードを終了します。
```

## ヘルプ

```shellsession
$ qiitrans --help
QiiTrans
  main コマンドは文書作成支援ツールです。
  標準入力のテキストを、引数の言語から言語へ自動翻訳した結果を返します。標準入力がない場合は、対話モードになります。

Usage:
  main [Options] LangFrom LangTo [LangTo ...]

Example:
  $ # 日本語 → 英語に翻訳
  $ echo '私は、賛成の反対に同意なのだ' | main ja en
  I agree to disagree.

  $ # 日本語 → 英語 → 日本語に翻訳
  $ echo '私は、賛成の反対に同意なのだ' | main ja en ja
  同意しないことに同意します。

  $ # 日本語 → 英語 → スペイン語 → 中国語 → 日本語の順に翻訳
  $ echo '同意しないことに同意します。' | main ja en es zh ja
  同意しないことに同意します。

LangFrom:
  標準入力から受け取るテキストの言語を指定します。（翻訳元の言語）

LangTo:
  翻訳先の言語を指定します。
  LangTo は複数指定可能で、1 つ前の言語から翻訳します。

翻訳エンジンについて:
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

Options:

  -a, --apikey           翻訳に使うエンジンのアクセス・トークンを指定します
      --cache-id         キャッシュの DB 名。異なる DB にキャッシュを保存したい場合に指定します
  -e, --engine[=deepl]   翻訳に使うエンジンを指定します
      --clear            実行前にキャッシュを完全に削除します。（API の利用枠を消費します）
  -h, --help             ヘルプを表示します
      --debug            デバッグ情報を標準エラー出力に出力します
      --no-cache         キャッシュを利用せずに翻訳 API から再取得します。（API の利用枠を消費します）
      --verbose          中間翻訳も出力します。
      --info             API のリクエスト可能な残数など、API 情報を出力します
  -v, --version          アプリのバージョン情報を表示します

```

## アクセス・トークン

`QiiTrans` は、翻訳 API のアクセストークン／認証キーを必要とします。環境変数もしくはコマンドの引数で指定してください。

- 現在は [DeepL API](https://www.deepl.com/docs-api/) のみ対応しています。[DeepL のアカウント](https://www.deepl.com/pro-account/summary)から認証キーを発行してください。

```shellsession
$ env | grep DEEPL_API_KEY
DEEPL_API_KEY=12345678-ffff-ffff-ffff-123456789abc:fx

$ echo '私は、賛成の反対に同意なのだ' | qiitrans ja en
...
```

```shellsession
$ export DEEPL_API_KEY=12345678-ffff-ffff-ffff-123456789abc:fx

$ echo '私は、賛成の反対に同意なのだ' | qiitrans ja en
...
```

```shellsession
$ echo '私は、賛成の反対に同意なのだ' | qiitrans --apikey "12345678-ffff-ffff-ffff-123456789abc:fx" ja en
...
```

## コントリビュート

- テストの実行にも DeepL のアクセストークンを環境設定にセットする必要があります。
- GitHub 上で `.`（ドット）キーを押すと Web エディターが起動します。
- [質問・要望・改善案・提案](https://github.com/Qithub-BOT/QiiTrans/discussions) @ Discussions
- [不具合報告（要再現テストあり）](https://github.com/Qithub-BOT/QiiTrans/issues) @ Issues
- [TODO](https://github.com/Qithub-BOT/QiiTrans/issues) @ Issues
  - [WIP](https://github.com/Qithub-BOT/QiiTrans/pulls) | Draft PR @ Pull Requests

## ライセンスと著作権

- License: [MIT](https://github.com/Qithub-BOT/QiiTrans/blob/master/LICENSE.md)
- Copyright: (c) [The QiiTrans Contributors](https://github.com/Qithub-BOT/QiiTrans/graphs/contributors)
