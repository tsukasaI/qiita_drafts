# GoとRustの文字列周りのデータ型

筆者は普段Goを使って開発していてデータ型に関して馴染があった。

最近はRustを勉強し始めて文字列周りのデータ型について気になったことがあったのでまとめていく。

## Go

Goの文字列にまつわるデータ型は以下の通り。

- string
- byte
- rune

### string

一般的に文字列を扱うためのデータ型。

特に補足はいらないと思うがコードで例を見てみる。

```go
s := "Hello, World!"
fmt.Println(s)

// "Hello, World!"
```

ダブルクオーテーションで囲まれた文字列がstring型として扱われる。

### byte

byteはuint8のエイリアスであり、バイトデータを扱うためのデータ型。

コードで例を見てみる。

```go

b := []byte("Hello, World!")
fmt.Println(b)

// [72 101 108 108 111 44 32 87 111 114 108 100 33]
```

このように文字列をbyte型に変換することでその文字のASCIIコードが出力される。

### rune

runeはint32のエイリアスであり、Unicode文字を扱うためのデータ型。

コードで例を見てみる。

```go

r := 'a'

fmt.Println(r)


// 97
```

Goではシングルクオーテーションで囲まれた文字がrune型として扱われる。

このようにUnicode文字をrune型に変換することでその文字のUnicodeコードポイントが出力される。

### それぞれの関係

string, rune, byteの関係を簡単に見るためにサンプルコードを書いた。

```go
package main

import (
	"fmt"
)

func main() {
	s := "a"
	fmt.Printf("s type is %T\n\n", s)

	r := 'a'
	fmt.Printf("r type id %T\n\n", r)

	fmt.Println("---for loop over string---")
	for i, c := range s {
		fmt.Printf("s[i] type is %T\n\n", s[i])
		fmt.Printf("c type is %T\n\n", c)
	}
}
```

実行すると以下のようになる。
```
str type is string

str[0] type is uint8

r type id int32

---for loop over string---
str[i] type is uint8

c type is int32
```

ここからわかることは以下の通り

- stringはuint8のスライスであり、その要素（インデックスでアクセスしたとき）はuint8型となっている
- rangeでstringをループすると、その要素はrune型となっている

## Rust

Rustの文字列にまつわるデータ型は以下の通り。

- str
- String
- char

### str

Rustの文字列リテラルはstr型として扱われる。

str型は不変であり、固定長の文字列を表す。

コードで例を見てみる。

```rust
fn main() {
    let s: &str = "Hello, World!";
    println!("{}", s);
}
```

ダブルクオーテーションで囲まれた文字列がstr型として扱われる。

### String

Stringは可変であり、ヒープ上に確保される可変長の文字列を表す。

Stringのstructはu8のバイトのベクタを持っている。

コードで例を見てみる。

```rust

fn main() {
    let s: String = String::from("Hello, World!");
    println!("{}", s);
}
```

String型はString::from関数を使って生成する。

### char

charはUnicodeスカラー値を表す。

コードで例を見てみる。

```rust

fn main() {
    let c: char = 'a';
    println!("{}", c);
}
```

シングルクオーテーションで囲まれた文字がchar型として扱われる。

## 最後に

文字に関するデータ型も複数あるのでその違いを理解しておくとプログラムを書くときに意識しましょう。
