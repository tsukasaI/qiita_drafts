# Go言語でホットリロードをしながら爆速でAPI開発をできる環境構築 ~Air, docekr-compose~

## はじめに

これまでGoの基礎の記事を公開してきた

https://www.ariseanalytics.com/activities/report/20221005/
https://www.ariseanalytics.com/activities/report/20220704/

この記事では実際にGoでAPIを開発するときに使うツールを紹介する。

## Goを用いる際の問題

過去の記事でGoはコンパイラ言語であると紹介しました。

APIサーバーを開発する際に問題になるのは「ソースコードの変更をする度にコンパイル操作をする必要がある」ことで開発効率を落としてしまいます。

## Air

そこで開発されたのがAirです。

https://github.com/cosmtrek/air

開発者はGoのAPIサーバーを構築時に即座にリロードがされないことに不満を感じホットリロードツールのAirを作りました。

本記事ではAirを用いたAPIサーバー開発環境の構築方法を紹介します。

## 使用する技術

Go v1.20
Gin v1.9.0 (GoのHttpフレームワーク、本記事での説明は割愛する)
Docker
docker-compose

## 本記事のゴール

docker-compose up コマンドを実行するとAPIサーバーが起動し
ソースコードを変更したときに自動で再ビルドして変更が反映される

# 各ファイルの紹介

ディレクトリ構成

```
.
├── Dockerfile
├── docker-compose.yaml
├── go.mod
├── go.sum
├── main.go
```

```Dockerfile
FROM golang:1.20.4-bullseye

RUN go install github.com/cosmtrek/air@latest
```

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

.air.toml（Airの設定ファイル）
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

Ginのドキュメントのサンプルコードをそのまま記述する。

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

```mod:go.mod
module air_sample

go 1.20

require github.com/gin-gonic/gin v1.9.0

require (
	github.com/bytedance/sonic v1.8.0 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.11.2 // indirect
	github.com/goccy/go-json v0.10.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.0.9 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.0.6 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.9 // indirect
	golang.org/x/arch v0.0.0-20210923205945-b76863e36670 // indirect
	golang.org/x/crypto v0.5.0 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
```

# 動作チェック

## httpサーバーが起動するか

```
$ docker compose up
[+] Running 1/0
 ✔ Container air_sample  Created                                     0.0s
Attaching to air_sample
air_sample  |
air_sample  |   __    _   ___
air_sample  |  / /\  | | | |_)
air_sample  | /_/--\ |_| |_| \_ , built with Go
air_sample  |
air_sample  | watching .
air_sample  | !exclude tmp
air_sample  | building...
air_sample  | running...
air_sample  | [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
air_sample  |
air_sample  | [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
air_sample  |  - using env:     export GIN_MODE=release
air_sample  |  - using code:    gin.SetMode(gin.ReleaseMode)
air_sample  |
air_sample  | [GIN-debug] GET    /ping                     --> main.main.func1 (3 handlers)
air_sample  | [GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
air_sample  | Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
air_sample  | [GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
air_sample  | [GIN-debug] Listening and serving HTTP on :8080
```

このようにHttpサーバーが起動しました。
実際にリクエストを実行するとレスポンスが返ってくることが確認できます。

```
$ curl localhost:8080/ping
{"message":"pong"}
```

dockerのターミナル
```
air_sample  | [GIN] 2023/05/10 - 04:35:23 | 200 |     367.667µs |      172.18.0.1 | GET      "/ping"
```

## ホットリロードが機能するか

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
に変更。入るを上書き保存をすると

dockerのターミナル
```
air_sample  | main.go has changed
air_sample  | building...
air_sample  | main.go has changed
air_sample  | running...
```

レスポンスの確認
```
$ curl localhost:8080/ping
{"message":"pong updated"}
```

このように保存をすると自動でビルドをし直してサーバーが起動するようになった。
