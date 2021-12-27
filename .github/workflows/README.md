# Workflow/Actions to Automate

- [merte-tests.yaml](merge-tests.yaml): マージ前に必要なテストを実行します。

## Fork 時の注意

Fork 先のリポジトリでも翻訳 API の動作テストのためにアクセストークン（認証キー）の設定が必要です。
リポジトリの `secrets` の設定でアクセストークン登録してください。

### Repository secrets の設定

Fork 先のリポジトリの [Settings] -> [Secrets] -> "Actions secrets" の "New repository secrets" ボタンから以下の設定でトークンを設定します.

- DeepL PRO API （[認証キー確認先](https://www.deepl.com/pro-account/plan)）
  - Name: `DEEPL_API_KEY`
  - Value: DEEPL で発行した認証キー
