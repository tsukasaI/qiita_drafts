# GoでRedisを使ったときの基本の備忘録

本記事は業務で利用したGoとRedisを使ったときの備忘録です。

AWSのElaticacheを使い、GoのRedisクライアントは`github.com/go-redis/redis/v8`を使うのが一般的のようだったのでその使い方をまとめる。

（Get、Setの基本のみ書くため、詳細は公式ドキュメントを参照してください）

## GoでRedisを使う

大まかな流れは

1. Redisクライアントを作成
1. Redisの操作

という超シンプルな流れです。

### Redisクライアントを作成

go-redis/redisのクライアントを作成するには、`NewClient`を使います。

```go
import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
)

func NewRedisClient() *redis.Client {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    return rdb
}

```

Redisのアドレス、パスワード、DBを指定して`NewClient`を使ってクライアントを作成します。

Optionsにはリトライ回数やタイムアウトと、プールサイズなども指定できます。

### Redisの操作

超基本の使い方のデータをセットする、データを取得するを書いてみます。

```go
func SetData(ctx context.Context, rdb *redis.Client, key string, value string) error {
    err := rdb.Set(ctx, key, value, 0).Err()
    if err != nil {
        return fmt.Errorf("failed to set data: %w", err)
    }

    return nil
}

func GetData(ctx context.Context, rdb *redis.Client, key string) (string, error) {
    val, err := rdb.Get(ctx, key).Result()
    if err != nil {
        return "", fmt.Errorf("failed to get data: %w", err)
    }

    return val, nil
}

// SetDataとGetDataの使い方

func ExecuteSetAndGet() {
    // Redisクライアントを作成
    ctx := context.Background()
    rdb := NewRedisClient

    // データをセット
    err := SetData(ctx, rdb, "key", "value")
    if err != nil {
        fmt.Println(err)
    }

    // データを取得
    val, err := GetData(ctx, rdb, "key")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(val)
}

```

ここでGetの戻り値の方について解説します。


上記の書き方をすると`GetData`関数の`rdb.Get`は[StringCmd型](https://pkg.go.dev/github.com/go-redis/redis/v8#StringCmd)を返します。
StringCmdのResultで返される第一返り値は`string`です。

一方でStringCmdからIntCmdやFloatCmdに変換して`int`や`float`などの型に変換することもできます。

```go
// int型を返す
func GetIntData(ctx context.Context, rdb *redis.Client, key string) (int, error) {
    val, err := rdb.Get(ctx, key).Int()
    if err != nil {
        return 0, fmt.Errorf("failed to get data: %w", err)
    }

    return val, nil
}

// float64型を返す
func GetFloatData(ctx context.Context, rdb *redis.Client, key string) (float64, error) {
    val, err := rdb.Get(ctx, key).Float64()
    if err != nil {
        return 0, fmt.Errorf("failed to get data: %w", err)
    }

    return val, nil
}

```

## 最後に

GoでRedisを使うときの基本の使い方をまとめました。

インメモリデータベースとして使うことでパフォーマンスを向上させることができるので、ぜひ使ってみてください。
