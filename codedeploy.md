# はじめに

本記事ではEC2、GitHub Actions、CodeDeployを用いたオートデプロイのやり方についてまとめる。

コンテナやサーバレスアーキテクチャが主流になりつつある現在でもオンプレミスでの運用を行っているところもまだまだあり、
追加のコストを払わずしてCI/CDの構築の方法の参考になれば幸いと思い公開します。

# GitHub Actionsとは

GitHub Actionsはビルド、テスト、デプロイのパイプラインを自動化できるCI/CDのプラットフォームでGitHubのリポジトリに対してワークフローを実行することができる。
ワークフローはyamlファイルで管理され、手動でのトリガーや特定のブランチにコミットがプッシュされたときをトリガーをきっかけに実行できる。

料金はパプリックリポジトリでは無料、プライベートリポジトリでは月当たり2000分までは無料で使用できる。

# CodeDeployとは

EC2, ECS, Lambdaなどにソフトウェアをデプロイすることができるフルマネージドサービスです。

機能としては
- デプロイの自動化
- 多数のホストへのデプロイ
- 正常性とロールバックをモニタリング
などがあり、使用言語に依存しないCD環境を作成できる。

デプロイの手順はappspec.yamlに記述することでシェルスクリプトの実行が可能。

料金は「AWS CodeDeploy を使用した Amazon EC2、AWS Lambda や Amazon ECS へのコードデプロイに追加料金は必要ありません。」とされているためEC2を使っていると無料で使用できる。


# 目標

developブランチにコミットがプッシュされたときにオートデプロイを実行する

## イメージ

（drawioのイメージ）

1. GutHub ActionsでdevelopブランチへのプッシュをトリガーにCodeDeployをコール
1. appspec.yamlに従いコマンドを実行

## Github Actions の設定

リポジトリの `Settings > Secrets > Actions` でクレデンシャルを保存。

## workflowの設定

```yaml:codedeploy.yaml
name: CodeDeploy Caller

on:
  push:
    branches:
      - develop

jobs:
  codedeploy:
    runs-on: ubuntu-latest

    steps:
      - name: Checkout repository
        uses: actions/checkout@v3

      - name: CodeDeploy
        env:
          AWS_ACCESS_KEY_ID: ${{ secrets.AWS_ACCESS_KEY_ID }}
          AWS_SECRET_ACCESS_KEY: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
          AWS_REGION: ap-northeast-1
        run: |
          aws deploy create-deployment \
            --application-name sample_app-app \
            --deployment-group-name sample_app_group \
            --github-location repository=${{github.repository}},commitId=${{github.sha}} \
            --region ap-northeast-1 \
            --file-exists-behavior OVERWRITE

```

## appspec.yamlの設定

```yaml:appspec.yaml
version: 0.0
os: linux
files:
  - source: /
    destination: /var/www/code_deploy_test
hooks:
  BeforeInstall:
    - location: scripts/clear_cache.sh
      timeout: 300
      runas: root
  AfterInstall:
    - location: scripts/start_app.sh
      timeout: 600
      runas: root
```

## shellの設定

```sh:scripts/start_app.sh
#!/bin/bash

# permissionの変更
chown -R ec2-user:ec2-user /var/www/code_deploy_test
chmod -R 777 /var/www/code_deploy_test/laravel/storage

# アプリケーションのスタート
cd /var/www/code_deploy_test/docker/production
docker-compose up -d
```

## CodeDeployの設定
```terraform:main.tf
# ----------------------------------
# IAM
# ----------------------------------
resource "aws_iam_role" "sample_app_deploy_role" {
  name = "sample_app_deploy_role-role"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "Principal": {
        "Service": "codedeploy.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "code_deploy_policy_attachments" {
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSCodeDeployRole"
  role       = aws_iam_role.sample_app_deploy_role.name
}

# ----------------------------------
# CodeDeploy
# ----------------------------------
resource "aws_codedeploy_app" "sample_app" {
  name = "sample_app-app"
}

resource "aws_codedeploy_deployment_group" "sample_app_deploy_group" {
  app_name              = aws_codedeploy_app.sample_app.name
  deployment_group_name = "sample_app_group"
  service_role_arn      = aws_iam_role.sample_app_deploy_role.arn

  ec2_tag_set {
    ec2_tag_filter {
      key   = "Name"
      type  = "KEY_AND_VALUE"
      value = "production"
    }
  }

  auto_rollback_configuration {
    enabled = true
    events  = ["DEPLOYMENT_FAILURE"]
  }
}
```

ec2_tag_set はEC2を識別するためのタグの条件で`Name = production`としているEC2に対してデプロイを行うようにしました。

GitHubのリポジトリでActionsタブを確認して、以下の記述があれば成功

```
Run aws deploy create-deployment \
{
    "deploymentId": "d-XXXXXXXXX"
}
```
