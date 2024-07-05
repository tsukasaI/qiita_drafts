# Goの空の構造体

Goの特定のデータセットを扱える構造体で、A Tour of Goにも出てくる。

```go
package main

import "fmt"

type Vertex struct {
    X int
    Y int
}

func main() {
    fmt.Println(Vertex{1, 2})
}
```

この例ではX, Yをintとして持つVertexという構造体を定義している。

そんな構造体のフィールドを持たない空の構造体を使うこともある。

空の構造体の特徴と使う例をまとめる。

## 空の構造体

空の構造体は以下のようなものである。

```go
type emptyStruct struct{}

var emptyStructVar struct{}{}
```

再度の説明にはなってしまうが、空の構造体はフィールドを持たない構造体である。

フィールドを持たないため使い道がないじゃん？と思うかもしれないが、メモリを消費しない。

通常の構造体はフィールドを持つためその分のメモリを確保するが、空の構造体はフィールドがないためメモリを消費しないという動きになる。

## 使い方

メモリを消費しないことを利用して以下のような使い方がある。

### セット型の表現

Pythonなどには実装されているセットをGoで表現するときに使う。

```go
func main() {
    set := make(map[string]struct{})
    set["a"] = struct{}{}
    set["b"] = struct{}{}
    set["c"] = struct{}{}

    // check if a exists
    if _, ok := set["a"]; ok {
        fmt.Println("a exists")
    }

    fmt.Println(set)
}
```

mapのvalueにboolを入れると少なからずメモリを消費するが、空の構造体を入れることでメモリを消費しない。

これとmapの2つ目の返り値を使ってkeyが存在するかどうかを判定することができSetのように使うことができる。

### チャネルのクローズ

チャネルのシグナル送信の際にも使用可能。

Goのchanはint, stringなどの型を指定するが、空の構造体を指定することでメモリを消費しない。

```go
func main() {
    ch := make(chan struct{})

    go func() {
        time.Sleep(1 * time.Second)
        close(ch)
    }()

    <-ch
    fmt.Println("done")
}
```

## まとめ

空の構造体はフィールドを持たない構造体でメモリを消費しない。

値に意味を持たせずにメモリ消費を抑えることが可能になる。
