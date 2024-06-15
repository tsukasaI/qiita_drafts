# curlでいろんな結果を見るならwオプションが便利だった

APIサーバーの開発をして簡単にレスポンスタイムなど性能を見たいありますよね。

毎回k6などでテストをするのも面倒だし、curlで簡単に見れないかなーと思って調べたところ、wオプションが便利だったのでメモしておきます。

## 結論

--write-out(-w)でtime_totalを表示すると、リクエストにかかった時間がわかります。

```bash
curl --request GET \
 -w"time_total: %{time_total}\n" \
 --url ＄URL
```

## write-outオプション

`-w`オプションを使うことで、curlの出力をカスタマイズできます。

サンプルとして`-w '%{response_code}\n' $URL`とすると、レスポンスコードが表示されます。


## よく使うフィールド

| フィールド | 説明 |
| ---- | ---- |
| %{http_code} | レスポンスコード |
| %{time_connect} | 接続にかかった時間 |
| %{time_appconnect} | アプリケーションに接続するのにかかった時間 |
| %{time_pretransfer} | リクエストが開始されてから最初のバイトが転送されるまでの時間 |
| %{time_starttransfer} | リクエストが開始されてから最初のバイトが転送されるまでの時間 |
| %{time_total} | リクエストにかかった合計時間 |


```bash

curl --request GET \
 -w"http_code: %{http_code}\ntime_connect: %{time_connect}\ntime_appconnect: %{time_appconnect}\ntime_pretransfer: %{time_pretransfer}\ntime_starttransfer: %{time_starttransfer}\ntime_total: %{time_total}\n" \
 --url $URL
```


参考:
https://github.com/curl/curl/blob/aab0c16990ff4cccbb68ab9fa9c81823b288fdcc/docs/cmdline-opts/write-out.md?plain=1#L4

https://qiita.com/yasushi-jp/items/bcc67b76d2491a1bc09e
（こちらの記事を参考にしてお仕事が進みました。ありがとうございました。）
