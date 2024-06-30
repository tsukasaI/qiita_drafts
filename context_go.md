# Goのcontext.Contextを感覚的に理解する

Goだけでなく様々な言語でContext, ctxが使われるので、自分の理解のためにまとめてみる。

## Context

まず言葉の定義から。

contextとは文脈や前後関係を意味します。

プログラミングの文脈では、関数の実行において必要な情報を渡すための仕組みである。

HTTPリクエストにはpath, method, header, bodyなど様々な情報が含まれているが、これらの情報をcontextが保持している。

## Goのcontextパッケージ

参考は[こちら](https://pkg.go.dev/context)。

ざくっとまとめるとこのパッケージはContextタイプを定義して、タイムアウトの共有やキャンセルの共有、リクエストスコープ値の共有などを提供している。

context.Contextはインターフェースで、以下のメソッドを持っている。

```go
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}
```

このContextはあるAPIの処理において、タイムアウトやキャンセルを行うために使われ、複数のgoroutine間で共有し各メソッドがコールされる。

## 使い方

context.Contextは`context.WithCancel`, `context.WithDeadline`, `context.WithTimeout`, `context.WithValue`などの関数を使って生成する。

基本的には`context.Background()`を使って生成するのがよいっぽい。

```go
ctx := context.Background()
```

このctxを関数の引数に渡すことで、関数内でContextのメソッドを使うことができる。

```go
func doSomething(ctx context.Context) {
    select {
    case <-ctx.Done():
        return
    default:
        // do something
    }
}
```

このようにgoroutine内でctx.Done()を監視することで、親のgoroutineがキャンセルされたときに子のgoroutineもキャンセルされる。

## まとめ

Contextを使うことで値の共有やキャンセルを簡単に実装できる。

ライブラリ
