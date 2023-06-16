Go言語でホットリロードをしながら爆速でAPI開発をできる環境構築 ~Air, Docker, docekr-compose~

はじめに

Marketing Solution Division所属のエンジニアの井上です。

これまでARISE tech blogでは基礎的なGo言語（以下Go）記事を公開してきました。

[【Go言語入門】「Go」の概要や特徴とは？社内のAPI開発プロジェクトで採用して分かったこと](https://www.ariseanalytics.com/activities/report/20221005/)

[【Go言語入門】goroutineとは？ 実際に手を動かしながら goroutineの基礎を理解しよう！](https://www.ariseanalytics.com/activities/report/20220704/)

今回はGoの応用編としてAPI開発をする際に使う便利なツールを紹介していきます。

サンプルのコードもありますのでぜひお楽しみください。

# GoでAPIを開発する際の問題

過去の記事でGoはコンパイラ言語であると紹介しました。

コンパイラ言語は一般的に実行時の処理速度が速く、Goはその速度からAPIサーバー開発言語に多くの企業で採用されています。

しかしAPIサーバーを開発する際に「ソースコードの変更をする度にコンパイル操作をする必要がある」

という問題を抱えていて開発効率を落としてしまいます。

そこで開発されたのが[Air](https://github.com/cosmtrek/air)です。

開発者はGoのAPIサーバーを構築時に即座にリロードがされないことに不満を感じホットリロードツールを作りました。

本記事ではそんなAirを用いたAPIサーバー開発環境の構築方法を紹介します。

# 本記事のゴール

docker-compose up コマンドを実行するとAPIサーバーが起動する。

ソースコードを変更したときに自動で再ビルドして変更が反映される。

サンプルのコードはこちら。
https://github.com/ariseanalytics/air_sample


# 構築

## 実行条件

Go v1.20

Gin v1.9.0 (GoのHttpフレームワーク、本記事での説明は割愛します)

Docker v23.0.5

docker-compose v2.17.3


## 各ファイルの紹介

### ディレクトリ構成

```
.
├── Dockerfile
├── docker-compose.yaml
├── go.mod
├── go.sum
├── main.go
```

### Dockerfile

```Dockerfile
FROM golang:1.20.4-bullseye

RUN go install github.com/cosmtrek/air@latest
```

コンテナにgoコマンドを利用してairをインストールします。

### docker-compose.yaml

```yaml:docker-compose.yaml
version: "3.8"

services:
  go:
    container_name: air_sample
    volumes:
      - ./:/project/
    working_dir: /project
    tty: true
    build: "./"
    ports:
      - 8080:8080
    command: sh -c 'go mod tidy && air'
```

コンテナを起動したらGoで用いるパッケージのインポートとAirの起動を行います。

### .air.toml

```toml:.air.toml
root = "."
testdata_dir = "testdata"
tmp_dir = "tmp"

[build]
  args_bin = []
  bin = "./tmp/main"
  cmd = "go build -o ./tmp/main ."
  delay = 0
  exclude_dir = ["assets", "tmp", "vendor", "testdata"]
  exclude_file = []
  exclude_regex = ["_test.go"]
  exclude_unchanged = false
  follow_symlink = false
  full_bin = ""
  include_dir = []
  include_ext = ["go", "tpl", "tmpl", "html"]
  include_file = []
  kill_delay = "0s"
  log = "build-errors.log"
  poll = false
  poll_interval = 0
  rerun = false
  rerun_delay = 500
  send_interrupt = false
  stop_on_error = false

[color]
  app = ""
  build = "yellow"
  main = "magenta"
  runner = "green"
  watcher = "cyan"

[log]
  main_only = false
  time = false

[misc]
  clean_on_exit = false

[screen]
  clear_on_rebuild = false
  keep_scroll = true
```

.air.tomlはAirの設定ファイルであり `air init` で生成したファイルをそのまま利用します。

簡単に説明をすると

- "_test" を含むファイルと"assets", "tmp", "vendor", "testdata" ディレクトリはビルドの対象外とする
- tmp/ ディレクトリにmainという実行ファイルをビルドして実行する

という設定にしています。


### main.go

```go:main.go
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
```

Ginのドキュメントのクイックスタートに記載のサンプルコードをそのまま記述しました。

このように記述すると8080番ポートでhttpリクエストを待ち受けるようになります。

# 動作チェック

## httpサーバーが起動するか

コンテナを起動するために`docker compose up`を実行します。

```
$ docker compose up
[+] Running 1/0
 ✔ Container air_sample  Created           0.0s
Attaching to air_sample
air_sample  |
air_sample  |   __    _   ___
air_sample  |  / /\  | | | |_)
air_sample  | /_/--\ |_| |_| \_ , built with Go
air_sample  |

~中略~

air_sample  | [GIN-debug] Listening and serving HTTP on :8080
```

このようにHttpサーバーが起動しました。
実際にリクエストを実行するとレスポンスが返ってくることが確認できます。

```
$ curl localhost:8080/ping
{"message":"pong"}
```

dockerのターミナルを確認すると以下のように出力されます。
```
air_sample  | [GIN] 2023/05/10 - 04:35:23 | 200 |     367.667µs |      172.18.0.1 | GET      "/ping"
```

## ホットリロードが機能するか

main.goを下記のように変更し上書き保存をしてみます。

```go:main.go
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong updated",
		})
	})
	r.Run()
}
```

保存をした時にdockerのターミナルを見ると
```
air_sample  | main.go has changed
building...
running...

~中略~

[GIN-debug] Listening and serving HTTP on :8000
```

と表示され再ビルドとAPIの再起動が行われました。

レスポンスの確認をしてみると

```
$ curl localhost:8080/ping
{"message":"pong updated"}
```

このように変更が反映されるようになりました。

# 最後に

今回はGoの応用事例としてAPIサーバーの開発に使う便利なツールを紹介しました。

我々の開発でも重宝しているツールなのでGoでAPIサーバーを構築する際にはぜひ使ってみてください。

最後まで読んでくださりありがとうございました。

今後も近年ホットな技術の記事もアップしていきますので見に来てください！
