# DistrolessというDockerイメージがよさげ

Googleが提供しているDockerイメージの一つにDistrolessというものがある。

GitHubリンクは[こちら](https://github.com/GoogleContainerTools/distroless)

本記事ではこのイメージの特徴と簡単な利用方法について説明する。

## 一般的なDockerイメージ

例えば、DebianやUbuntuなどのOSイメージを使う場合、OSに関連するパッケージやツールが含まれています

しかし実際にアプリケーションを商用利用する場合、これらのパッケージやツールは不要であることが多く、不要なパッケージやツールが含まれています。

利用しないパッケージやツールが含まれていることで生じる問題は以下の通りです。

- セキュリティリスク
- イメージサイズの増加
- メンテナンスの複雑化

以上から、アプリケーションの実行に必要な最低限のコンポーネントのみを含むDockerイメージがあると便利です。

## Distrolessの特徴

一方で、Distrolessはアプリケーションの実行に必要な最低限のコンポーネントのみを含むDockerイメージです。

上記で記載した内容の裏返しにはなりますが、Distrolessの特徴は以下の通りです。

- OSに関連するパッケージやツールが含まれていない
- イメージサイズが小さい
- セキュリティリスクが低い

これにより、アプリケーションの実行に必要な最低限のコンポーネントのみを含むDockerイメージを利用することで、セキュリティリスクを低減し、イメージサイズを小さくすることができます。

これを商用環境には特に有用です。

## Distrolessの利用方法

筆者がこのDockerイメージを推す理由はただひとつ。

**GoでビルドしたバイナリをDockerイメージに含めるだけで、実行環境を構築できるから**です。

Goは静的バイナリを生成することができるため、DistrolessのDockerイメージに含めるだけで実行環境を構築が可能です。

CPUアーキテクチャの差に気をつけてビルドする必要がありますが、上記のメリットを享受できるためGoでアプリケーションを開発する際には本番環境にDistrolessのDockerイメージを利用することを個人的にはおすすめします。

## 簡単に試してみる

簡単にDockerfileを作ってみます。

今回は単純にホストマシン上のテキストファイルをコピーしてビルドするだけのDockerfileを作成します。

```Dockerfile
# Distroless Static Dockerfile
FROM gcr.io/distroless/static:nonroot

#  Copy an text file
COPY sample.txt /sample.txt
```

sample.txtは以下の内容です。

```txt
Hello from distroless container
```

このイメージをビルドしてみます。

```sh
% distroless % docker build -t test-distroless .

~~~

% docker images
REPOSITORY                   TAG                            IMAGE ID       CREATED         SIZE
test-distroless              latest                         xxxxxxxxxxxx   2 minutes ago   1.99MB
```

イメージサイズが1.99MBと非常に小さくなっていることが確認できます。

さて、このコンテナでcatコマンドを実行してみようとします。

```sh
% docker run --rm test-distroless cat /sample.txt

docker: Error response from daemon: failed to create task for container: failed to create shim task: OCI runtime create failed: runc create failed: unable to start container process: exec: "cat": executable file not found in $PATH: unknown.
```

このように、DistrolessのDockerイメージにはcatコマンドなどのシェルが含まれていないため、コンテナ内でシェルを実行することができません。

ここでGoでビルドしたバイナリをDistrolessのDockerイメージに含めることでコマンドを実行することができます。

```go
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	cat("/sample.txt")
}

func cat(filePath string) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(file))
}
```

このコードをビルドしてDockerfileに含めることで、DistrolessのDockerイメージに含めることができます。

```Dockerfile
# Distroless Static Dockerfile
FROM gcr.io/distroless/static

#  Copy an text file
COPY sample.txt /sample.txt

#  Run the application
COPY main /main
```

```sh
% distroless % go build -o main main.go
% distroless % docker build -t test-distroless .

~~~

% docker run --rm test-distroless
Hello from distroless container
```

このように、GoでビルドしたバイナリをDistrolessのDockerイメージに含めることで、コンテナ内でコマンドを実行することができることが確認できます。
