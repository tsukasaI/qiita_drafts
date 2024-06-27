# qiita_drafts

## future

1. algorithmの勉強でやったソートをGoでまとめる
1. orverload
1. Rustの基本勉強
    1. 構造体とかimplとか
1. スケールするWebアプリケーションの設計
    1. キャッシュの基本
1. Goにおけるキャッシュの使われどころ調べる
1. Goの1.22のnet/http ServeMuxがよさげ
1. プロジェクトのディレクトリ構成でよくある形と意味
1. たまに使うgitコマンド（辞書みたいに使う）
1. こんなコマンドあったんだ git log --graph
1. AWSの構成図とかER図はDraw.ioをVSCodeで作成してGitHubで管理しませんか
1. pklというAppleのオープンソースの設定ファイルを見てみた https://pkl-lang.org/index.html
1. Goのテストにおけるdeferとt.Cleanupの違い
1. コードファーストとスキーマファーストどう採用するか
1. エラーハンドリングの良しあしについての個人的意見
1. 個人的によく使うVSCodeのショートカット
1. Golangの空の構造体 https://zenn.dev/mkosakana/articles/e24f4a3aeb3027
1. VScodeのショートカットを極限まで使いこなす
1. context.Contextを感覚的に理解する
1. DynamoDBをローカルで動かしてデータを入れたり見たりする
1. Dockerのコンテナ、イメージ、ボリューム、キャッシュなどを一括で削除する
- Goでmockを自動生成

docker run -v "$PWD":/src -w /src vektra/mockery --all

```
  private getKeys<T extends { [key: string]: unknown }>(obj: T): (keyof T)[] {
    return Object.keys(obj);
  }
```

- GoのCIでテストのカバレッジをコメントするようにしたった
- Gomockを使ってテストのmock入れたら意外と大変だった
- WinとMacで部分的にスクショを撮るショートカット
  - Win
  - Mac
- copilotは万能ではない話
  - 全ておまかせではこっちの意図は汲み取ってくれない。（あくまで学習して予想したサジェストをしてくれるだけ）
- GoでHeap使える話
- クイックソート全然理解できない
