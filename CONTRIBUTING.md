# コントリビュート（参加）について

## PullRequest 先

- `main` ブランチ
- PR に慣れていない方 👉 [How To Pull Request](https://github.com/Qithub-BOT/QiiCipher/blob/master/.github/HOW_TO_PULL_REQUEST.md) | QiiCipher | Qithub @ GitHub

## マージ・ルール

- `squash` `merge` マージ

テストおよびレビューをパスした PR は、コミットを 1 つにまとめてマージします。そのため、**コミット・ログの綺麗さは気にする必要はありません**。どんどん試行錯誤してください。

## PR 時の注意（レギュレーション）

PR もしくは [Draft PR](https://github.blog/jp/2019-02-19-introducing-draft-pull-requests/) されると CI による各種テストが自動実行されます。

これらのテストをパスしないと、レビューおよびマージは行われません。

1. 構文・文法チェック（lint）と静的解析
    - Golang 及びシェルスクリプトの構文テストが行われます。
        - `shfmt` ........... シェル・スクリプトの構文チェック
        - `shellcheck` ...... シェル・スクリプトの静的解析
        - `gofmt` ........... Golang の構文チェック
        - `golangci-lint` ... Golang の静的解析チェック
2. ユニット・テスト
    - `go test ./src/...` によるユニット・テストが行われます。`./.github/run-tests-coverage.sh` でもテストできます。
3. カバレッジ・テスト
    - 各々のパッケージ（ディレクトリ）にテストが 1 つでも設置されていた場合、カバレッジが 100% にならないとパスしません。
    - 未カバーの箇所を確認したい場合は ``./.github/run-tests-coverage.sh --verbose` で確認できます。
4. ドキュメントの更新チェック
    - 各パッケージ（ディレクトリ）にある README.md は [gomarkdoc](https://github.com/princjef/gomarkdoc) で自動作成されたドキュメントです。
    - CI では変更がドキュメントに反映されているかのテストが行われます。`./.github/update-docs.sh` で更新してください。(要 gomarkdoc)

これらの確認や開発に必要な環境の Dockerfile を [./.devcontainer](.devcontainer) に用意しています。ご活用ください。

## なんかわけわかめ

- 遠慮なく [Discussions](https://github.com/Qithub-BOT/QiiTrans/discussions) の Q&A で聞いてください。
