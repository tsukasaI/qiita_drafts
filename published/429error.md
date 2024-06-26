# 429 Too Many Requests: その理解と実装

スケールを考慮したシステムでは一般的にリクエストの制限を行うことがあります。その中でも、特定の時間内に許可されるリクエスト数を制限するRate Limitingは、サービスの安定性やセキュリティを確保するために重要な機能です。

Rate Limitingとは、特定の時間内にユーザーやシステムがリソースにアクセスできる回数を制限することです。これは、リソースの過剰な使用を防ぎ、サービスの公平な使用を保証し、サービスの乱用を防ぐために使用されます。

HTTPプロトコルでは、Rate Limitingは主にHTTPステータスコードを通じて通知され、特に`429 Too Many Requests`は、ユーザーが許可されたリクエストの上限を超えたときに返されるステータスコードです。

429エラーが返されると、通常はRetry-Afterヘッダも含まれます。このヘッダは、クライアントが次にリクエストを再試行すべき時間を秒数または日付で示します。

したがって、429エラーを受け取った場合、クライアントはRetry-Afterヘッダの指示に従ってリクエストを再試行する、もしくはリクエストの頻度を減らす必要があります。

## Rate Limitingのアルゴリズム

Rate Limitingのアルゴリズムには主に以下の3つがあります。

- 固定窓アルゴリズム (Fixed Window Algorithm): このアルゴリズムでは、一定の時間窓（ex. 1分、1時間）ごとにリクエストの上限を設定する。時間窓が切り替わるとリクエストのカウントはリセットされる。しかし、時間窓の切り替わり直前に大量のリクエストが来るとその瞬間だけサーバーに大きな負荷がかかる可能性があｒｙ。

- スライディングウィンドウアルゴリズム (Sliding Window Algorithm): このアルゴリズムは固定窓アルゴリズムの問題を解決できる。スライディングウィンドウアルゴリズムでは、過去の一定時間内（ex. 過去1分間、過去1時間）のリクエスト数をカウントする。これにより時間窓の切り替わり直前の大量のリクエストによるサーバーへの負荷を緩和できる。

- トークンバケットアルゴリズム (Token Bucket Algorithm): このアルゴリズムでは、一定の速度でトークンがバケットに追加され、各リクエストはトークンを消費する。バケットが空になると、新たなリクエストはトークンが追加されるまで待たされます。このアルゴリズムは一定の速度でリクエストを処理しつつ、短期的なバーストを許容することができる。

これらのアルゴリズムはそれぞれ異なるシナリオに適しています。固定窓アルゴリズムは実装が簡単ですがスライディングウィンドウアルゴリズムやトークンバケットアルゴリズムの方が、リクエストの流量をより滑らかに制御できます。

## トークンバケットアルゴリズムの実装

トークンバケットアルゴリズムをGoで実装する例を示します。

```go
package main

import (
    "fmt"
    "time"
)

type TokenBucket struct {
    capacity int
    tokens   int
    rate     int
    lastTime time.Time
}

func NewTokenBucket(capacity, rate int) *TokenBucket {
    return &TokenBucket{
        capacity: capacity,
        tokens:   capacity,
        rate:     rate,
        lastTime: time.Now(),
    }
}

func (tb *TokenBucket) Allow() bool {
    now := time.Now()
    elapsed := now.Sub(tb.lastTime)
    tb.lastTime = now

    // Add tokens
    tb.tokens += int(elapsed.Seconds()) * tb.rate
    if tb.tokens > tb.capacity {
        tb.tokens = tb.capacity
    }

    // Check if request is allowed
    if tb.tokens > 0 {
        tb.tokens--
        return true
    }

    return false
}

func main() {
    tb := NewTokenBucket(10, 1)

    for i := 0; i < 15; i++ {
        if tb.Allow() {
            fmt.Println("Request allowed")
        } else {
            fmt.Println("Request denied")
            time.Sleep(time.Second)
        }
    }
}
```

この実装では、トークンバケットの容量と速度を指定して、リクエストが許可されるかどうかを判定します。

上記のサンプルコードでは、トークンバケットの容量を10、速度を1として、15回のリクエストを処理しています。

トークンバケットが空の場合はリクエストが拒否され、トークンが追加される速度でトークンが補充されます。

## Redisを使用したRate Limiting

しかし複数のサーバーでRate Limitingを行う場合、各サーバーでトークンバケットアルゴリズムを独立して実装するとリクエストの流量が均等に分散されない可能性があります。

そのためトークンバケットアルゴリズムを実装する際に、Redisなどのインメモリデータベースを使用することで分散環境でのRate Limitingを実現することができます。

Redisは高速でスケーラブルなキャッシュやデータストアとして広く利用されており、トークンバケットアルゴリズムの実装に適しています。

トークンバケットの状態をRedisに保存し、各リクエストごとにトークンの消費と補充を行うことで、分散環境でのRate Limitingを実現できます。

Redisを使用したRate Limitingの実装例は以下のようになります。

```go
package main

import (
    "fmt"
    "time"

    "github.com/go-redis/redis/v8"
)

type RateLimiter struct {
    client *redis.Client
    key    string
    limit  int
    period time.Duration
}

func NewRateLimiter(client *redis.Client, key string, limit int, period time.Duration) *RateLimiter {
    return &RateLimiter{
        client: client,
        key:    key,
        limit:  limit,
        period: period,
    }
}

func (rl *RateLimiter) Allow() bool {
    now := time.Now()
    pipe := rl.client.TxPipeline()
    pipe.ZRemRangeByScore(rl.client.Context(), rl.key, "-inf", now.Add(-rl.period).Format(time.RFC3339))
    pipe.ZCard(rl.client.Context(), rl.key)
    pipe.ZAdd(rl.client.Context(), rl.key, &redis.Z{Score: float64(now.Unix()), Member: now.Unix()})
    pipe.Expire(rl.client.Context(), rl.key, rl.period)
    cmders, err := pipe.Exec(rl.client.Context())
    if err != nil {
        return false
    }
    count, _ := cmders[1].(*redis.IntCmd).Result()
    return count <= int64(rl.limit)
}

func main() {
    client := redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
    })
    defer client.Close()

    rl := NewRateLimiter(client, "rate_limit", 10, time.Minute)

    for i := 0; i < 15; i++ {
        if rl.Allow() {
            fmt.Println("Request allowed")
        } else {
            fmt.Println("Request denied")
        }
    }
}
```

この実装では、RedisのZSET（ソート済み集合）を使用して、トークンバケットの状態を保存し、リクエストの許可を判定しています。
リクエストが許可されるかどうかは、ZSETに保存されたタイムスタンプを元に判定されます。

ロジックの具体のステップは以下のようになっています。

1. 現在の時間を取得: now := time.Now()で現在の時間を取得します。

1. トランザクションパイプラインを開始: pipe := rl.client.TxPipeline()でRedisのトランザクションパイプラインを開始します。これにより、複数のRedisコマンドを一度に実行することができます。

1. 古いリクエストを削除: pipe.ZRemRangeByScore(rl.client.Context(), rl.key, "-inf", now.Add(-rl.period).Format(time.RFC3339))で、現在の時間から設定した期間を引いた時間より前のリクエストを削除します。

1. 現在のリクエスト数を取得: pipe.ZCard(rl.client.Context(), rl.key)で、現在のリクエスト数を取得します。

1. 新たなリクエストを追加: pipe.ZAdd(rl.client.Context(), rl.key, &redis.Z{Score: float64(now.Unix()), Member: now.Unix()})で、新たなリクエストを追加します。リクエストのスコアとメンバーには現在のUnix時間を設定します。

1. キーの有効期限を設定: pipe.Expire(rl.client.Context(), rl.key, rl.period)で、キーの有効期限を設定します。有効期限は設定した期間になります。

1. トランザクションを実行: cmders, err := pipe.Exec(rl.client.Context())で、上記のすべてのコマンドを一度に実行します。

1. エラーチェック: if err != nil { return false }で、トランザクションの実行中にエラーが発生したかどうかをチェックします。エラーが発生した場合は、新たなリクエストは許可されません。

1. リクエスト数をチェック: count, _ := cmders[1].(*redis.IntCmd).Result()で、現在のリクエスト数を取得します。return count <= int64(rl.limit)で、リクエスト数が設定した上限以下であるかどうかをチェックします。上限以下であれば新たなリクエストは許可され、上限を超えていれば許可されません。

このように、Redisを使用したRate Limitingの実装により分散環境でのリクエストの流量を制御し、サービスの安定性を確保することができます。
