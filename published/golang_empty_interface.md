# Goの空のインターフェースであらゆる型を代入できるのってなんで？

Goの空のインターフェースは、どんな型でも受け取ることができるため、あらゆる型を代入できます。

```go
package main

import "fmt"

func main() {
    var anyVal interface{}
    anyVal = 1
    fmt.Println(anyVal) // 1

    anyVal = "string"
    fmt.Println(anyVal) // string

    anyVal = []int{1, 2, 3}
    fmt.Println(anyVal) // [1 2 3]

    anyVal = map[string]int{"key": 1}
    fmt.Println(anyVal) // map[key:1]
}
```

これがなぜできるのかを知ったときになるほどと思ったため説明していきます。

## インターフェース

そもそもインターフェースとはなんでしょうか？

[こちらの記事](https://qiita.com/tsukasaI/items/7b6516adcf85bb96e249)で説明していますが、インターフェースはメソッドの集まりを示すものになります。

ある型に対してその型が持つメソッドの集まりを定義することができ、具体的な処理の中身を来にすること無く、その型が持つメソッドを呼び出すことができます。

例えば、以下のようなインターフェースがあるとします。

```go
type Stringer interface {
    String() string
}
```

このインターフェースは`String`メソッドを持つ型に対して適用されます。

```go
package main

import "fmt"

type MyString string

func (m MyString) String() string {
    return string(m)
}

func main() {
    var s Stringer
    s = MyString("Hello, World")
    fmt.Println(s.String()) // Hello, World
}
```

このStringerはMyString型以外にも適用することができ、一例としては数値型にも適用することができます。

```go
type StringerInt int

func (s StringerInt) String() string {
    return fmt.Sprintf("%d", s)
}
```

このようにすることでStringerInt型もStringerインターフェースを満たすことができます。

また、インターフェースで期待していないメソッド以外も持っていても問題ありません。

```go
type StringerInt64 int64

// Stringerインターフェースを満たす
func (s StringerInt64) String() string {
    return fmt.Sprintf("%d", s)
}

func (s StringerInt64) Int() int64 {
    return s
}
```

このようにStringerInt64型はStringerインターフェースを満たしているため、Stringer型に代入することができます。

## さて本題

ここまでインターフェースについて説明してきましたが、空のインターフェースはどうでしょうか？

空のインターフェースには満たすべきメソッドが存在しないという状態になります。

つまりどんなメソッドを持っていても、もしくは何もメソッドを持っていなくても空のインターフェースに代入することができます。

このことからint型、string型、slice型、map型、それ以外の任意のstruct型など、どんな型でも代入することができるということになります。

Go 1.18からはanyが導入され、interface{}のaliasとなっています。

## 最後に

（初めてinterface{}が一般的なanyになるか説明できる？と言われたときに説明できなかったのでこの記事で言語化しておきました。）

みなさんはドヤ顔で説明できるようになりましょう！
