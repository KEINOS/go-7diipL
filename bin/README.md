# 実行バイナリ用ディレクトリ

`build-app.sh` スクリプトを実行すると、各種 OS（アーキテクチャ）用にバイナリを作成し、zip アーカイブします。

ローカルに Go がインストールされていない場合でも Docker を検知した場合は Alpine の軽量イメージでバイナリのビルドを行います。

```bash
# 対応アーキテクチャ一覧表示
./build-app.sh --list

# ヘルプ
./build-app.sh --help

# Linux AMD64/Intel アーキテクチャ向けにビルド
./build-app.sh linux

# macOS AMD64/Intel アーキテクチャ向けにビルド
./build-app.sh darwin

# Windows AMD64/Intel アーキテクチャ向けにビルド
./build-app.sh windows

# RaspberryPi OS ARM v6 アーキテクチャ向けにビルド
./build-app.sh linux arm 6

# 上記の一般的なアーキテクチャを一括ビルド
./build-app.sh --all
```
