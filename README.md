# qiita_drafts

## future

1. algorithmの勉強でやったソートをGoでまとめる
1. orverload
1. イベントソースの例と実装
1. Rustの基本勉強
    1. メモリ周り
    1. 構造体とかimplとか
1. ファイルの最終は改行を入れる
1. GoとRedis使った備忘録
1. スケールするWebアプリケーションの設計
    1. キャッシュの基本
1. Goにおけるキャッシュの使われどころ調べる
1. 構造的にAPI設計を考える
1. Goの1.22のnet/http ServeMuxがよさげ
1. プロジェクトのディレクトリ構成でよくある形と意味
1. たまに使うgitコマンド（辞書みたいに使う）
1. こんなコマンドあったんだ git log --graph
1. Larevelで名言を表示する php artisan inspire が誰得だった
1. AWSの構成図とかER図はDraw.ioをVSCodeで作成してGitHubで管理しませんか
1. pklというAppleのオープンソースの設定ファイルを見てみた https://pkl-lang.org/index.html
1. Goのテストにおけるdeferとt.Cleanupの違い
1. コードファーストとスキーマファーストどう採用するか

1. curlでいろんな結果を見るならwオプションが便利だった

curl --request GET \
 -w"http_code: %{http_code}\ntime_namelookup: %{time_namelookup}\ntime_connect: %{time_connect}\ntime_appconnect: %{time_appconnect}\ntime_pretransfer: %{time_pretransfer}\ntime_starttransfer: %{time_starttransfer}\ntime_total: %{time_total}\n" -o /dev/null \
 --url https://dev-api.wellness.auone.jp/v2/presents

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
