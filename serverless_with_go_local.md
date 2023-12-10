Serverless Framework 導入 ローカル環境でGoをLambdaで実行する例

# Serverlessめっちゃ便利

ローカルでLambdaの実行環境を構築する場合に使う


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

```
% npm run dev

> go_lambda_serverless@1.0.0 dev
> serverless offline

(node:26841) [DEP0040] DeprecationWarning: The `punycode` module is deprecated. Please use a userland alternative instead.
(Use `node --trace-deprecation ...` to show where the warning was created)

Starting Offline at stage api (ap-northeast-1)

Offline [http for lambda] listening on http://localhost:3002
Function names exposed for local invocation by aws-sdk:
           * api: aws-lambda-go-api-proxy-gin-api-api

   ┌───────────────────────────────────────────────────────────────────────┐
   │                                                                       │
   │   GET | http://localhost:3000/api/ping                                │
   │   POST | http://localhost:3000/2015-03-31/functions/api/invocations   │
   │   GET | http://localhost:3000/api/hello                               │
   │   POST | http://localhost:3000/2015-03-31/functions/api/invocations   │
   │                                                                       │
   └───────────────────────────────────────────────────────────────────────┘

Server ready: http://localhost:3000 🚀


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



GET /api/hello (λ: api)
✖ START RequestId: a077c5a7-0586-1eb3-42b0-edcd500bcf99 Version: $LATEST

✖ [GIN] 2023/12/09 - 12:09:55 | 404 |         667ns |                 | GET      "/hello"

✖ END RequestId: a077c5a7-0586-1eb3-42b0-edcd500bcf99

✖ REPORT RequestId: a077c5a7-0586-1eb3-42b0-edcd500bcf99        Duration: 8.96 ms  Billed Duration: 9 ms    Memory Size: 1024 MB    Max Memory Used: 63 MB

✖ Handler/layer file changed, restarting bootstrap...
  Handler/layer file changed, restarting bootstrap...
  Handler/layer file changed, restarting bootstrap...
  Handler/layer file changed, restarting bootstrap...



GET /api/hello (λ: api)
✖ 2023/12/09 12:10:32 Gin cold start

✖ [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

  [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
   - using env: export GIN_MODE=release
   - using code:        gin.SetMode(gin.ReleaseMode)

  [GIN-debug] GET    /ping                     --> main.init.0.func1 (3 handlers)
  [GIN-debug] GET    /hello                    --> main.init.0.func2 (3 handlers)

✖ START RequestId: 98f2c140-0ae9-1cfc-bc0c-9878c62d773e Version: $LATEST

✖ [GIN] 2023/12/09 - 12:10:32 | 200 |      19.416µs |                 | GET      "/hello"

✖ END RequestId: 98f2c140-0ae9-1cfc-bc0c-9878c62d773e

✖ REPORT RequestId: 98f2c140-0ae9-1cfc-bc0c-9878c62d773e        Duration: 5.38 ms  Billed Duration: 6 ms    Memory Size: 1024 MB    Max Memory Used: 63 MB
```
