Serverless Framework 導入 ローカル環境でGoをLambdaで実行する例

# Serverlessめっちゃ便利

ローカルでLambdaの実行環境を構築する場合に使うツールにServerless Frameworkがあります。

いつもローカル環境ではDocker、サーバーはECSなりEC2などを使っていたところからマイクロサービスを構築できるように勉強したので

それを共有します。

## 今回のお題

- Goを使ってAPIを構築する
- パッケージはGinを使う
- APIをLambdaを使ってコールできるようにする

これらを満たすやり方を紹介していきます。

## 環境

- M2 Macbook Air
- Go v1.21.4 darwin/arm64
- npm 10.2.4
- Docker 24.0.6, build ed223bc

お急ぎの方は[こちら](https://github.com/tsukasaI/serverless_with_go)から

## コードたち


### Go

まずはGoでAPIを構築していきましょう。

GinとLambdaを使うのでそれぞれインポートします。

今回はついでにenvの読み取り方も示すのでenvパッケージも使います。

```go:main.go
package main

import (
	"context"
	"log"

	env "github.com/caarlos0/env/v10"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

type config struct {
	SAMPLE_ENV string `env:"SAMPLE_ENV"`
}

var ginLambda *ginadapter.GinLambda

func init() {
	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Gin cold start")

	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("%+v\n", err)
	}
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "world",
		})
	})

	ginLambda = ginadapter.New(r)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}

```

エントリーポイントであるmain関数にはlambda.Startをコールして、Handler関数を呼び出します。

mainよりも先にinit関数が呼ばれて各種設定を行います。

envの設定についてはconfig構造体を定義して、env.Parseをして値を入れます。

SAMPLE_ENVについてはserverless.yamlで定義するので後ほど。

`r := gin.Default()`以降はでルーティングを行います。

シンプルにするためにエンドポイントのレスポンスは固定の文字列にしました。


設定が終わったらginとLambdaをつなぐために`ginLambda = ginadapter.New(r)`としてHandlerで利用します。

### serverless.yaml

```yaml:serverless.yaml
service: aws-lambda-go-api-proxy-gin

provider:
  name: aws
  architecture: arm64
  environment:
    SAMPLE_ENV: SAMPLE_ENV
  runtime: provided.al2
  stage: ${opt:stage, self:custom.defaultStage}
  region: ap-northeast-1
  iam:
    role:
      statements:
        - Effect: "Allow"
          Action:
            - "logs:*"
          Resource: "*"

plugins:
  - serverless-go-plugin
  - serverless-offline

custom:
  defaultStage: api
  serverless-offline:
    useDocker: true

package:
  individually: true
  exclude:
    - "./**"

functions:
  api:
    handler: bootstrap
    timeout: 100
    events:
      - http:
          path: ping
          method: get
      - http:
          path: hello
          method: get
```

ポイントは

- ENVの設定
ENVはprovider > environmentで定義します。

今回は`SAMPLE_ENV: SAMPLE_ENV`としてセットしました。

ここはyamlファイルに載せていい内容にする、もしくはParameter Storeなどのサービスを使って設定するようにしましょう。


- Runtimeの設定
provider > runtimeで設定します。

ここで注意ですが、go1.xのランタイムはサポートが終了するので`provided.al2`を使いましょう。

こちらはビルドしたバイナリが必要になるので別途ビルドできるようにする必要があります。

- serverless-offlineのオプション
provided.al2を利用する際にはdockerを使用します。

custom > serverless-offline > useDockerをtrueにしておきましょう

## ローカル環境で実行

ビルドしてserverless offlineコマンドを実行してみよう。

（参考のGitHubにはnpmでビルドとserverless offlineコマンドは実行できるようにしてあります。）

```
% serverless offline

> go_lambda_serverless@1.0.0 dev
> serverless offline

~中略~

   ┌───────────────────────────────────────────────────────────────────────┐
   │                                                                       │
   │   GET | http://localhost:3000/api/ping                                │
   │   POST | http://localhost:3000/2015-03-31/functions/api/invocations   │
   │   GET | http://localhost:3000/api/hello                               │
   │   POST | http://localhost:3000/2015-03-31/functions/api/invocations   │
   │                                                                       │
   └───────────────────────────────────────────────────────────────────────┘

Server ready: http://localhost:3000 🚀

```

これで起動できました。

`curl http://localhost:3000/api/ping`でリクエストしてみると

```
GET /api/ping (λ: api)
✖ Lambda API listening on port 9001...

✖ 2023/12/09 12:09:51 Gin cold start

✖ [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

  [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
   - using env: export GIN_MODE=release
   - using code:        gin.SetMode(gin.ReleaseMode)

  [GIN-debug] GET    /ping                     --> main.init.0.func1 (3 handlers)

✖ START RequestId: 01490924-e0e4-1335-a097-7dd6f3cfdc51 Version: $LATEST

✖ [GIN] 2023/12/09 - 12:09:51 | 200 |      17.208µs |                 | GET      "/ping"

✖ END RequestId: 01490924-e0e4-1335-a097-7dd6f3cfdc51

✖ REPORT RequestId: 01490924-e0e4-1335-a097-7dd6f3cfdc51        Init Duration: 94.07 ms     Duration: 12.88 ms      Billed Duration: 13 ms  Memory Size: 1024 MB    Max Memory Used: 61 MB
```

のように表示され、レスポンスは`{"message":"pong"}`となっています。

みんなもServerlessでGoのAPIを動かしてみてください。
