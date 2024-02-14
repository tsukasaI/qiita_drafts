# qiita_drafts

## future

- algorithmの勉強でやったソートをGoでまとめる
- orverload

```
  private getKeys<T extends { [key: string]: unknown }>(obj: T): (keyof T)[] {
    return Object.keys(obj);
  }
```

* GoのCIでテストのカバレッジをコメントするようにしたった
* Gomockを使ってテストのmock入れたら意外と大変だった
* WinとMacで部分的にスクショを撮るショートカット
    * Win
    * Mac
* copilotは万能ではない話
    * 全ておまかせではこっちの意図は汲み取ってくれない。（あくまで学習して予想したサジェストをしてくれるだけ）
* ファイルの最終は改行を入れる
* GoでHeap使える話
* クイックソート全然理解できない
- curlでいろんな結果を見るならwオプションが便利だった

curl --request GET \
   -w"http_code: %{http_code}\ntime_namelookup: %{time_namelookup}\ntime_connect: %{time_connect}\ntime_appconnect: %{time_appconnect}\ntime_pretransfer: %{time_pretransfer}\ntime_starttransfer: %{time_starttransfer}\ntime_total: %{time_total}\n" -o /dev/null \
  --url https://dev-api.wellness.auone.jp/v2/presents

- Goでmockを自動生成

docker run -v "$PWD":/src -w /src vektra/mockery --all
