# Goのbuildtagをサクッとリードしてみる。

builtin.goはGolangのリポジトリの`src/builtin/builtin.go`に書かれている。

このファイルにはGoで使われる定数、type、関数が定義されている。

当たり前に使っている定数や関数がどのように定義されているかをいくつか見てみましょう。


## bool

Goにおけるbool型は`true`と`false`の2つの値を持つ。


それぞれは`true`と`false`という定数で定義されている。


```go
const (
    true  = 0 == 0 // Untyped bool
    false = 0 != 0 // Untyped bool
)
```

（VSCodeのGoの拡張機能で`true`と`false`にホバーするとこの定義に飛べて、初めて見たときは驚いた）

## make

makeは以下のように定義されている。

```go
func make(t Type, size ...IntegerType) Type
```

makeはslice、map、channelに対しするメモリの確保を行う。

Typeがsliceの場合、sizeはlenとcapの両方を指定する。

```go
make([]int, 10, 100)
```

とすると、10個の要素を持つスライスを作成し、そのスライスの容量は100になる。

Typeがmapの場合はsizeはcapを指定する。

```go
make(map[string]int, 100)
```

とすると、100個の容量を持つマップを作成する。

Typeがchannelの場合はsizeはバッファの容量を指定する。

```go
make(chan int, 10)
```

とすると、10個のバッファを持つチャネルを作成する。


## new

newは以下のように定義されている。

```go
func new(Type) *Type
```

newはmakeとは異なり、Typeのゼロ値のポインタを返す。

## panic

panicは以下のように定義されている。

```go
func panic(v any)
```

panicはプログラムを停止させ、関数内でdeferされた関数を実行し、その後プログラムを終了させる。

panicをコールした関数がを呼び出す関数にもdeferされた関数があれば、それも実行される。

プログラムの終了時にはスタックトレースが表示され、0ではない終了コードが返される。

## recover

recoverは以下のように定義されている。

```go
func recover() any
```

recoverはpanicを呼び出した関数内でのみ使用できる。

recoverはpanicをコールした関数内でdeferされた関数内でのみ使用できる。

deferされていない関数の外でrecoverがコールされた場合panicはキャンセルされず、プログラムは停止する。panicしなかった場合はnilを返す。
