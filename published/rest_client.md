# はじめに

Web/モバイルエンジニアのみなさん

REST APIをデバッグなどでリクエストするときにはどんなツールを使っていますか？

本記事では筆者がこれまで使ってきたツールの比較を通してどのツールが良いかサクッと比較できるように書きます

## そもそもREST APIとは

詳細は割愛しますが、RESTの思想に沿った設計のAPIのこと。

アプリケーションで基本となるCRUDに対してリソース名とメソッドを組み合わせて表現することが多い。

HTTPリクエストのデバッグを行うためにリクエストを実行するツールが必要になる。

## ツール一蘭

筆者が使用経験のあるのは以下。

- curl
- Advanced REST client
- Postman
- REST Client

### curl

公式ページ: https://curl.se/

言わずと知れたコマンドで、URLに対してリクエストを実行できるオープンソースソフトウェア。

インストールしたらコマンドラインでリクエストが実行できる。

```sh

$ curl http://localhost

```

上記コマンドでGETリクエストが発行できる。

POSTを実行したい場合は-X or --requestオプションを追加する。

bodyのデータは-d or --data、ヘッダーは-H or --headerを追加する。


```sh

$ curl http://localhost \
    -H 'content-type: application/json' \
    -X PUT \
    -d '{"foo": "bar"}'

```

CLIですぐに試せるのは非常にありがたい一方でbodyやheaderが長くなるとコマンドが長くなり書きたくなくなってきますね。

### Advanced REST client

Google Chromeの拡張機能で、Chromeのストアから入手可能。
（2022年12月にサポートが終了したようです）

UIはかなりわかりやすいと思っていますがURLとメソッドを指定してリクエストします。

ボディとヘッダーも専用のタブがあるので選択して指定することができました。

### Postman

公式サイト: https://www.postman.com/

Advanced REST clientと同様にUIがわかりやすく、ヘッダーやボディをそれぞれ指定可能。

フォルダ階層で保存することも可能で、エンドポイント毎に保存することで複数のプロジェクトのリクエストを分けたりすることができます。

さらにcurlに出力することも可能で、`</>`マークをクリックするとcURLに現在のリクエストの内容をコマンドとして表示してくれます。

### REST Client

vscodeのサイト: https://marketplace.visualstudio.com/items?itemName=humao.rest-client

Postmanとは異なりコードエディタで準備することが必要

Usageにも書かれているが、URLを記載して`cmd + option + R`でGETリクエストが実行できる。

```http
https://example.com/comments/1
```

POSTでヘッダーとボディを追加する場合は以下のように先頭にメソッド名を書く。

```http
POST https://example.com/comments
Content-Type: application/json

{
    "foo": "bar"
}
```

UIは無いがファイルとして保存することでGitで共有が可能になり、プルリクエスト時に「ファイルに記載のリクエストを実行をする」という使い方ができる。


## どれが良いか

Advanced REST client以外で比較すると以下のようになる

- curl
  - CLIで実行可能
  - オプションで指定する必要があるためパラメータが多い場合はコマンドが長くなる
- Postman
  - UIがありわかりやすい
  - カーソル操作をする必要があるので、ウィンドウの切り替えが必要になる
- VSCode REST Client
  - VSCodeだけで完結して、ファイル保存してGitで共有可能
  - UIが無いため専用のコードに慣れる必要がある

個人的にはPostmanを使っていましたが、VSCodeで作業をまとめたいこととGitでリクエストを管理できるためVSCode REST Clientを積極的に使っています。
