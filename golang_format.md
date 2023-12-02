# Goのfmt.Printf()をpkg.go.devを見て勉強する

Goの書式付きフォーマットを見ると最初はびっくりしませんか？

過去の私と同じ経験をしているあなたと一緒にこの記事で勉強していきたいと思います。

参考サイト
https://pkg.go.dev/fmt


## 忙しい方向けにサマリ

| フォーマット | 出力 |
| -- | -- |
| %v | 通常のフォーマット |
| %+v | フィールド名つきのフォーマット |
| %d | int |
| %T | (型情報) |
| %t | bool |


* デバッグ時には%+vを進めたい

## 参考サイトを見ていきましょう

```go
package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	// A basic set of examples showing that %v is the default format, in this
	// case decimal for integers, which can be explicitly requested with %d;
	// the output is just what Println generates.
	integer := 23
	// Each of these prints "23" (without the quotes).
	fmt.Println(integer)
	fmt.Printf("%v\n", integer)
	fmt.Printf("%d\n", integer)

```

`%v` 通常のフォーマット、つまりどの型でもそのまま表示する形式です。

23というinteger型は明示的には `%d` で表示することを求められます。



```go
	// The special verb %T shows the type of an item rather than its value.
	fmt.Printf("%T %T\n", integer, &integer)
	// Result: int *int
```
`%T` は値自体ではなく対象のの型を表示します。

intergerと&intergerはint型、intのポインタ型と表示してくれます。


```go
	// Println(x) is the same as Printf("%v\n", x) so we will use only Printf
	// in the following examples. Each one demonstrates how to format values of
	// a particular type, such as integers or strings. We start each format
	// string with %v to show the default output and follow that with one or
	// more custom formats.

```

よく使うfmt.Println()はfmt.Printf("%v\n", n)と同じことを示します。

普段Printlnを使っている方は`%v`の結果を見ているのですね。


```go
	// Booleans print as "true" or "false" with %v or %t.
	truth := true
	fmt.Printf("%v %t\n", truth, truth)
	// Result: true true
```

bool型に対しては`%v` と `%t` で表示可能です。

```go
	// Integers print as decimals with %v and %d,
	// or in hex with %x, octal with %o, or binary with %b.
	answer := 42
	fmt.Printf("%v %d %x %o %b\n", answer, answer, answer, answer, answer)
	// Result: 42 42 2a 52 101010

	// Floats have multiple formats: %v and %g print a compact representation,
	// while %f prints a decimal point and %e uses exponential notation. The
	// format %6.2f used here shows how to set the width and precision to
	// control the appearance of a floating-point value. In this instance, 6 is
	// the total width of the printed text for the value (note the extra spaces
	// in the output) and 2 is the number of decimal places to show.
	pi := math.Pi
	fmt.Printf("%v %g %.2f (%6.2f) %e\n", pi, pi, pi, pi, pi)
	// Result: 3.141592653589793 3.141592653589793 3.14 (  3.14) 3.141593e+00

	// Complex numbers format as parenthesized pairs of floats, with an 'i'
	// after the imaginary part.
	point := 110.7 + 22.5i
	fmt.Printf("%v %g %.2f %.2e\n", point, point, point, point)
	// Result: (110.7+22.5i) (110.7+22.5i) (110.70+22.50i) (1.11e+02+2.25e+01i)
```

int型に対しては`%v` と `%d` で10進数表示可能です。

hex値は`%x`, 八進数は`%o`, 二進数は`%b`で表示可能です。

（普段Goで開発している筆者も知らなかったです。）

小数は`%g`で表示可能で、桁数を指定する場合は`%{桁数}.{小数点以下何位までか}f`を使います。

指数表示は`%e`で表示可能です。

虚数は`%g`, `%f`, `%e`で表示可能で、虚数部については`i`が追加されます。

```go
	// Runes are integers but when printed with %c show the character with that
	// Unicode value. The %q verb shows them as quoted characters, %U as a
	// hex Unicode code point, and %#U as both a code point and a quoted
	// printable form if the rune is printable.
	smile := '😀'
	fmt.Printf("%v %d %c %q %U %#U\n", smile, smile, smile, smile, smile, smile)
	// Result: 128512 128512 😀 '😀' U+1F600 U+1F600 '😀'
```

runeは基本intになるが、`%c`でUnicode文字として表示可能です。

`%q`とすることでシングルクォーテーションの中に文字を表示してくれます。

`%U`でUnicodeのhex値を出してくれ、また`%#U`でUnicodeのhex値とruneの文字列としての値を表示してくれます。。


```go
	// Strings are formatted with %v and %s as-is, with %q as quoted strings,
	// and %#q as backquoted strings.
	placeholders := `foo "bar"`
	fmt.Printf("%v %s %q %#q\n", placeholders, placeholders, placeholders, placeholders)
	// Result: foo "bar" foo "bar" "foo \"bar\"" `foo "bar"`
```

string型に対しては`%v` と `%s` で表示可能です。

`%q`でシングルクォーテーションで囲った状態、`%#q`でバッククォーテーションで囲った状態で表示可能です。

```go
	// Maps formatted with %v show keys and values in their default formats.
	// The %#v form (the # is called a "flag" in this context) shows the map in
	// the Go source format. Maps are printed in a consistent order, sorted
	// by the values of the keys.
	isLegume := map[string]bool{
		"peanut":    true,
		"dachshund": false,
	}
	fmt.Printf("%v %#v\n", isLegume, isLegume)
	// Result: map[dachshund:false peanut:true] map[string]bool{"dachshund":false, "peanut":true}
```

mapに対しては`%v`でキーと値、`%#v` でGoで定義した情報とともに表示可能です。

ちなみにMapはキーでソートしてから表示されます。

```go
	// Structs formatted with %v show field values in their default formats.
	// The %+v form shows the fields by name, while %#v formats the struct in
	// Go source format.
	person := struct {
		Name string
		Age  int
	}{"Kim", 22}
	fmt.Printf("%v %+v %#v\n", person, person, person)
	// Result: {Kim 22} {Name:Kim Age:22} struct { Name string; Age int }{Name:"Kim", Age:22}
```

structに対しては`%v`で各フィールドの値のみ、`%+v` でフィールド名付き、`%#v`でGoで定義したstructも表示されます。

```go
	// The default format for a pointer shows the underlying value preceded by
	// an ampersand. The %p verb prints the pointer value in hex. We use a
	// typed nil for the argument to %p here because the value of any non-nil
	// pointer would change from run to run; run the commented-out Printf
	// call yourself to see.
	pointer := &person
	fmt.Printf("%v %p\n", pointer, (*int)(nil))
	// Result: &{Kim 22} 0x0
	// fmt.Printf("%v %p\n", pointer, pointer)
	// Result: &{Kim 22} 0x010203 // See comment above.
```

ポインタに対しては`%v` では `&` を先頭につけて変数の値を表示し、

`%p` でhex値でポインタの値で、nilの場合は0x0が表示されます。

```go
	// Arrays and slices are formatted by applying the format to each element.
	greats := [5]string{"Kitano", "Kobayashi", "Kurosawa", "Miyazaki", "Ozu"}
	fmt.Printf("%v %q\n", greats, greats)
	// Result: [Kitano Kobayashi Kurosawa Miyazaki Ozu] ["Kitano" "Kobayashi" "Kurosawa" "Miyazaki" "Ozu"]

	kGreats := greats[:3]
	fmt.Printf("%v %q %#v\n", kGreats, kGreats, kGreats)
	// Result: [Kitano Kobayashi Kurosawa] ["Kitano" "Kobayashi" "Kurosawa"] []string{"Kitano", "Kobayashi", "Kurosawa"}
```

配列とスライスに対しては`%v` と `%q` で各要素の文字列として表示可能です。

```go
	// Byte slices are special. Integer verbs like %d print the elements in
	// that format. The %s and %q forms treat the slice like a string. The %x
	// verb has a special form with the space flag that puts a space between
	// the bytes.
	cmd := []byte("a⌘")
	fmt.Printf("%v %d %s %q %x % x\n", cmd, cmd, cmd, cmd, cmd, cmd)
	// Result: [97 226 140 152] [97 226 140 152] a⌘ "a⌘" 61e28c98 61 e2 8c 98
```

byte(uint8)のスライスに対しては`%v` と `%d` で数値の表示が、

`%s`, `%q` で文字列表示が、

`%x`, `% x` でhexの表示が可能になります。

```go
	// Types that implement Stringer are printed the same as strings. Because
	// Stringers return a string, we can print them using a string-specific
	// verb such as %q.
	now := time.Unix(123456789, 0).UTC() // time.Time implements fmt.Stringer.
	fmt.Printf("%v %q\n", now, now)
	// Result: 1973-11-29 21:33:09 +0000 UTC "1973-11-29 21:33:09 +0000 UTC"

}

```

Stringerを実装した型では文字列と同じように表示可能です。
