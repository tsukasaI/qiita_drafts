# qiita_drafts

## future

1. asdfでバージョン管理 -> miseでバージョン管理 + direnvの置き換え
1. コンテナ使う？ホストマシン使う？
1. 脱・Dokcer Desktop colimaでDockerを動かす
1. go.workでworkspace管理ができるってよ
1. OIDC/CIBAとか認証認可周りの歴史と動向
1. App2App for Auth
1. Introduction Pre-commit tool Lefthook
1. Secret lint
1. 冪等性についてまとめる
1. test container がテストに使えそうな
1. モジュラーモノリスの考え方
1. GQL Federation https://netflixtechblog.com/how-netflix-scales-its-api-with-graphql-federation-part-1-ae3557c187e2
1. connect-goが便利っぽい
1. Golang Functional Options Pattern https://golang.cafe/blog/golang-functional-options-pattern.html
1. Kubernetes入門
1. Google Cloud Professional Architectの勉強

1. jsのformatter, linterはbiomeが早くていい感じ
1. algorithmの勉強でやったソートをGoでまとめる
1. orverload
1. Rustの基本勉強
    1. 構造体とかimplとか
1. Goにおけるキャッシュの使われどころ調べる
1. Goの1.22のnet/http ServeMuxがよさげ
1. プロジェクトのディレクトリ構成でよくある形と意味
1. たまに使うgitコマンド（辞書みたいに使う）
1. こんなコマンドあったんだ git log --graph
1. AWSの構成図とかER図はDraw.ioをVSCodeで作成してGitHubで管理しませんか
1. pklというAppleのオープンソースの設定ファイルを見てみた https://pkl-lang.org/index.html
1. エラーハンドリングの良しあしについての個人的意見
1. context.Contextを感覚的に理解する
1. DynamoDBをローカルで動かしてデータを入れたり見たりする
1. TypeScriptでデコレータを定義
1. プロジェクトの事始め ~ ルール、初期構築、ドキュメント管理、コミュニケーションフロー ~
1. あってほしいドキュメント、なくてもいいドキュメント
1. Goのテストのカバレッジとか見られるからやってみよう
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
