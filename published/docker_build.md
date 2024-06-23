# Dockerのビルドを復習する

Dockerの基本をおさらいするためにイメージのビルドについて説明します。

## イメージとコンテナ

コンテナを実行するためには、イメージをビルドする必要があります。

イメージは、コンテナの実行に必要なファイルや設定をまとめたものでイメージはDockerfileというファイルにビルドの手順を記述することで作成できます。

```Dockerfile
FROM ubuntu

CMD ["echo", "Hello, Docker!"]
```

このファイルでは、Ubuntuのイメージをベースにして、`echo "Hello, Docker!"`を実行するイメージを作成しています。

## イメージのビルド

イメージをビルドするには、`docker build`コマンドを使います。

```sh
docker build -t my-image .
```

このコマンドは、カレントディレクトリにあるDockerfileを使ってイメージをビルドし、`my-image`という名前でタグ付けします。

## イメージの実行

イメージを実行するには、`docker run`コマンドを使います。

```sh
docker run my-image
```

このコマンドは、`my-image`という名前のイメージを実行します。

## レイヤー

イメージは、複数のレイヤーから構成されています。

レイヤーはFROM, RUN, COPY, ADD, CMDなどのDockerfileの命令によって追加され、変更差分を保持しています。


## キャッシュ

buildの高速化のために、Dockerはビルドの各ステップの結果をキャッシュしています。

Dockerfileの命令が変更されない限りキャッシュされた結果を再利用することができ、ビルドの高速化につながります。

## マルチステージビルド

ビルドのステップを分割することに最終的なイメージに不要なファイルや設定を含めないようにすることができます。

これをマルチステージビルドと呼び、Dockerfileに複数のFROM命令を使って複数のステージを定義することで実現できます。

例えばGoのアプリケーションをビルドする場合、バイナリをビルドするステージと実行するためのイメージを作成するステージを分けることができます。

サンプルは次に記載します。

### イメージサイズの比較

ここからはGoのアプリケーションをビルドする例を使って、マルチステージビルドの効果を確認します。

対象のGoのアプリケーションは以下のようなものとします。

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, playground")
}
```

```go.mod
module main

go 1.22.4
```

このアプリケーションをビルドするためのDockerfileは以下のようにします。

```Dockerfile
FROM golang:1.22
WORKDIR /work/app
COPY main.go go.mod ./
RUN go mod tidy
RUN go build .
CMD ["./main"]
```

このDockerfileを使ってイメージをビルドし、サイズを確認します。

```sh
$ docekr build . -t normal_build_sample
$ docker images

REPOSITORY          ~ SIZE
normal_build_sample ~ 859MB
```

では、マルチステージビルドを使ってイメージをビルドし、サイズを確認します。

Goのコードは同一としてDistrolessイメージを使ってマルチステージビルドを行います。

```Dockerfile
FROM golang:1.22 AS builder
WORKDIR /work/app
COPY main.go go.mod ./
RUN go mod tidy
RUN go build .

FROM gcr.io/distroless/static:nonroot
COPY --from=builder /work/app ./
CMD ["./main"]
```

この例では、Goのアプリケーションをビルドするステージと実行するためのイメージを作成するステージを分けています。

バイナリを実行するイメージには、Distrolessという軽量なイメージを使っています。

Distrolessイメージについては[こちら](https://qiita.com/tsukasaI/items/34302c89b3136be1b5c9)で解説していますのでぜひ参考にしてください。

```sh
$ docker build . -t multi_stage_build_sample
$ docker images
REPOSITORY               SIZE
multi_stage_build_sample 3.93MB
```

このように、マルチステージビルドを使うことで、不要なファイルや設定を含めないようにすることができ、イメージのサイズを小さくすることができます。
