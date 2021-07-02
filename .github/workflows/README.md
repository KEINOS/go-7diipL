# Workflow/Actions to Automate

- [merte-tests.yaml](merge-tests.yaml): マージ前に必要なテストを実行します。

## Fork 時の注意

Fork 先のリポジトリの設定で、翻訳 API のアクセストークン（認証キー）を設定・登録する必要があります。

### Repository secrets の設定

Fork 先のリポジトリの [Settings] -> [Secrets] -> "Actions secrets" の "New repository secrets" ボタンから以下の設定でトークンを設定します.

- DeepL PRO API （[認証キー確認先](https://www.deepl.com/pro-account/plan)）
  - Name: `DEEPL_API_KEY`
  - Value: DEEPL で発行した認証キー
