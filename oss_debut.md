# はじめに

OSSに初コントリビュートをした感想と簡単な流れについて記述する。

---

# OSSとは

オープンソースソフトウェアのこと

ソースコードの改変、再配布を自由に行なって良いもの

PHP, Java, Rubyなど超有名プログラミング言語もOSS

---

# 具体的に何やったん？

色々頑張ったけど要約すると以下が効果あった

- やり方をググる
- [good first issue](https://goodfirstissue.dev)をdigった
- すぐにリポジトリをforkして変更作業着手

---

## やり方をググる

**OSS コミット やり方** で検索してみた

だいたい以下のステップ

- コントリビュートできそうなissueを報告しているOSSをGitHubで探す
  - [good first issue](https://goodfirstissue.dev/)
  - [Github Help Wanted](http://github-help-wanted.com/)
  - [GitHub](https://github.com/)
- issueを見つけたら「私がやります！」宣言をする
- 作業をして取り込み依頼を出す

---

## good first issue とは
GitHubのissueのラベルのうち `good first issue` が含まれるものをまとめてくれているページ
good first issueは`初めてのコントリビュートに向いている Issue につけられる`ラベル

---

## issueを選び方
もちろん何でもコントリビュートできる訳ではない
ケイパや興味で選ぶ必要がある

個人的には
- 得意なのはGo, TypeScript, DevOps(GitHub Actions)
- 興味があるのはテスト、スクレイピング、ORM、HTTPフレームワーク

---

## 変更作業
私がやった流れは以下

1. とりあえず`Can I take this? :)`と投稿する
1. コントリビュートのルールを確認する (Readmeを読む)
1. 間髪を入れずリポジトリをfork（自分のアカウント配下にコピー）する
1. ブランチを切って変更作業をする
1. リモートリポジトリにpushしてコントリビュート先にプルリクエストを出す

---

# 実際にコントリビュートしたものたち
## Goのテストコードの修正
https://github.com/MontFerret/ferret/pull/781
https://github.com/gavv/httpexpect/pull/347

## GitHub Actions Warningの修正
https://github.com/fairhive-labs/go-pixelart/pull/16

## 不要な引数の削除
https://github.com/apache/camel-k/pull/4155
なんとapacheのリポジトリにコントリビュート

---

# 最後に
知らない人に自分のやったことを認めてもらうのは嬉しい
やろうぜOSSコントリビュート
