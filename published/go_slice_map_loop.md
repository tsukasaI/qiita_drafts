# GoのSliceとMapをforループで回すときのnilの扱いを動かしながら調べた

Goのfor rangeでSliceやMapをループ処理を行うことがたくさんあると思います。

rangeの対象のlengthが0やnilの場合、どのような挙動をするのかが気になったのでサンプルコードを書いて動かしてみました。

## Slice

釈迦に説法かもですが、Sliceとは可変長の配列のことで全ての要素の型は同じです。

```go
	ints := []int{1, 2, 3, 4, 5}
	fmt.Println("[]int{1, 2, 3, 4, 5}のlen:", len(ints))
	fmt.Println("[]int{1, 2, 3, 4, 5}のrange")
	for i, v := range ints {
		fmt.Println(i, v)
	}
```

こんなコードを書いたときにlenは5、for文は5回繰り返します。

アウトプットは以下のようになります。

```
[]int{1, 2, 3, 4, 5}のlen: 5
[]int{1, 2, 3, 4, 5}のrange
0 1
1 2
2 3
3 4
4 5
```

次からSliceがnilだったり、lengthが0だったりしたときの挙動を見ていきます。

### nilの場合

さて、以下のコードがあった時にはどうなるでしょうか。

```go
	var nilSlice []int
	fmt.Println("nilSliceのlen:", len(nilSlice))
	fmt.Println("nilSliceのrange")
	for i, v := range nilSlice {
		fmt.Println(i, v)
	}
```

ここで補足ですがnilSliceの形は[]intですが、中身はnilです。

アウトプットは以下のようになります。

```
nilSliceのlen: 0
nilSliceのrange
```

ここでのポイントはrangeにnilが入っているときはエラーにならずにfor文が一度も実行されないということです。

### lengthが0の場合

次にlengthが0の場合を見ていきます。

```go
	emptySlice := []int{}
	fmt.Println("[]int{}のlen:", len(emptySlice))
	fmt.Println("[]int{}のrange")
	for i, v := range emptySlice {
		fmt.Println(i, v)
	}
```

こちらに関しては自明かもしれませんが、lengthが0の場合はfor文が一度も実行されません。

アウトプットは以下のようになります。

```
[]int{}のlen: 0
[]int{}のrange
```

## Map

次にMapについて見ていきます。

基本ですがMapのループのサンプルを示すために以下のコードを見てみます。

```go
	kv := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("map[string]int{\"a\": 1, \"b\": 2, \"c\": 3}のlen:", len(kv))
	fmt.Println("map[string]int{\"a\": 1, \"b\": 2, \"c\": 3}のrange")
	for k, v := range kv {
		fmt.Println(k, v)
	}
```

こちらを実行すると以下のようなアウトプットが得られます。

```
map[string]int{"a": 1, "b": 2, "c": 3}のlen: 3
map[string]int{"a": 1, "b": 2, "c": 3}のrange
a 1
b 2
c 3
```

### nil/lengthが0の場合

こちらもSliceと同様に以下のコードを考えます。

（Sliceと似ているのでまとめて書きます）

```go
	var nilMap map[string]int
	fmt.Println("nilMapのlen:", len(nilMap))
	fmt.Println("nilMapのrange")
	for k, v := range nilMap {
		fmt.Println(k, v)
	}

	emptyMap := map[string]int{}
	fmt.Println("map[string]int{}のlen:", len(emptyMap))
	fmt.Println("map[string]int{}のrange")
	for k, v := range emptyMap {
		fmt.Println(k, v)
	}
```

こちらを実行すると以下のようになります。


```
nilMapのlen: 0
nilMapのrange
map[string]int{}のlen: 0
map[string]int{}のrange
```

## まとめ

nilのSlice/Mapに対してrangeを使うとエラーとならずにfor文が一度も実行されないという動きになります。

ある関数やメソッドの返り値がnilの可能性がある場合でも、rangeを使う前にはnilチェックを行わなくてもエラーにならないので、その点は便利だと感じました。
