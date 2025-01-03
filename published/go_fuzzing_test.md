# Fuzz testingとGo

プログラムのテストとして多くの手法がある。

効果的なテストの手法を選択することでシステムの安全性や信頼性を飛躍的に高めることができる。

本記事ではテスト手法の一つである**Fuzz testing**について紹介しGoを用いたコードで実験してみる。

## Fuzz testingとは

Fuzz testingとはソフトウェアのテストの手法の一つで、想定されているデータだけでなく想定されていないデータや不正なデータをランダムに生成して入力として与えて実行するテストである。

テスト対象はブラックボックス（つまり実装の中身は不明）として扱い、**自動的**に**大量の**異常データを入力して動作をチェックするというニュアンスとなる。

もしテスト中にプログラムが不適切な動作をしたりクラッシュをしたりした場合はリリース前に気づくことができるため品質の向上に寄与できる。

## Goでの例

説明だけではイメージしづらいと思うので実際にGoのコードを見てみましょう。

GoにはtesingパッケージにFuzz testingを実行するための[F struct](https://pkg.go.dev/testing#F)が定義されている。

これを用いたFuzz testingを実行してみる。

### サンプルプログラム

Goのfuzz testingは以下の型の引数のみ受け入れる関係上、byteのスライスを用いた例を提示する。

> - string, []byte
> - int, int8, int16, int32/rune, int64
> - uint, uint8/byte, uint16, uint32, uint64
> - float32, float64
> - bool

コードはbyteのスライスを比較して一致しているかを判定する自前で実装した。

```go: compare.go
func IsEuqalIntSlice(a, b []byte) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
```

ひと目見てわかるかもしれないがこのコードは不十分である。

例えばaのlenがbよりも長い場合はpanicを発生させるし、bがaよりも長い場合かつbの先頭からaのlenまでが一致していた場合はtrueを返す。

### ユニットテスト

これを通常のユニットテストでチェックするには以下のようなコードになる。

```go: compare_test.go
func TestIsEuqalIntSlice(t *testing.T) {
	tests := []struct {
		name   string
		a, b   []byte
		expect bool
	}{
		{
			name:   "same strings",
			a:      []byte("abcd"),
			b:      []byte("abcd"),
			expect: true,
		},
		{
			name:   "same len but doesn't match",
			a:      []byte("aaa"),
			b:      []byte("AAA"),
			expect: false,
		},
		{
			name:   "a is shorter",
			a:      []byte("aa"),
			b:      []byte("aaa"),
			expect: false,
		},
		{
			name:   "a is longer",
			a:      []byte("aaa"),
			b:      []byte("aa"),
			expect: false,
		},
	}
	for _, tc := range tests {
		actual := IsEuqalIntSlice(tc.a, tc.b)
		assert.Equal(t, tc.expect, actual)
	}
}
```

testsには4種のケースの組み合わせを確認できるようにしてある。

この中でnameが"a is shorter", "a is longer"はexpectとIsEuqalIntSliceの返り値が異なるためユニットテストは失敗する。

しかしこの2つのケースが書けなかった場合はテストは成功してしまう。

このように引数に無限の組み合わせがある場合にはテストケースが漏れている可能性がある。

### Fuzz testing

ではこれをFuzz testingで書いてみる。

```go: compare_test.go
func FuzzIsEuqalIntSlice(f *testing.F) {
	f.Fuzz(func(t *testing.T, a, b []byte) {
		_ = IsEuqalIntSlice(a, b)
	})
}
```

これを実行するには `go test -fuzz=Fuzz -fuzztime 10s`のようなコマンドを実行する。

なおGoではfuzz testingを時間指定無しで実行すると無限に続けるので`-fuzztime`オプションで時間指定すると良い。

```sh-session
% go test -fuzz=Fuzz -fuzztime 10s
fuzz: elapsed: 0s, gathering baseline coverage: 0/3 completed
fuzz: elapsed: 0s, gathering baseline coverage: 3/3 completed, now fuzzing with 8 workers
fuzz: minimizing 62-byte failing input file
fuzz: elapsed: 0s, minimizing
--- FAIL: FuzzIsEuqalIntSlice (0.02s)
--- FAIL: FuzzIsEuqalIntSlice (0.00s)
testing.go:1591: panic: runtime error: index out of range [1] with length 1
goroutine 82 [running]:
runtime/debug.Stack()

~略~

Failing input written to testdata/fuzz/FuzzIsEuqalIntSlice/a6f774463b71fa5c
To re-run:
go test -run=FuzzIsEuqalIntSlice/a6f774463b71fa5c
```

このようにpanicで終了し、更にテストに使ったデータをファイルで保存してくれている。

```:testdata/fuzz/FuzzIsEuqalIntSlice/a6f774463b71fa5c
go test fuzz v1
[]byte("\xe30")
[]byte("\xe3")

```

再実行のコマンドが表示されているので、バグを修正したらこのコマンドを実行すると同じデータで実行することができる。


参考までにコードは以下のように修正するとテストがパスする。

```go:compare.go
func IsEuqalIntSlice(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
```

またFuzzingにもassertを追加する。

ちなみにf.Addを追加することで任意のテストケースの組み合わせをFuzzingにも追加することができる。

```go:compare_test.go
func FuzzIsEuqalIntSlice(f *testing.F) {
	cases := []struct{ a, b string }{
		{
			a: "abcd",
			b: "abcd",
		},
		{
			a: "aa",
			b: "aaa",
		},
	}
	for _, v := range cases {
		f.Add([]byte(v.a), []byte(v.b))
	}
	f.Fuzz(func(t *testing.T, a, b []byte) {
		expected := reflect.DeepEqual(a, b)
		actual := IsEuqalIntSlice(a, b)

		assert.Equalf(t, expected, actual, "Doesn't match. a: %+v, b: %+v", a, b)
		_ = IsEuqalIntSlice(a, b)
	})
}

```

```sh-session
% go test -run=FuzzIsEuqalIntSlice/a6f774463b71fa5c
PASS
ok  	qiita	0.256s
% go test . -run TestIsEuqalIntSlice
ok  	qiita	0.217s
```

## 最後に

プログラム/システムの品質向上には効果的なテストが不可欠である。

そのためには人間の手で書くコードのみならず自動で入力を設定するFuzz testingは予測不可能な入力があり得る場面で効果を発揮する。

GoのコードのテストでもjsonのMarshal, Unmarshalで使われている。

https://github.com/golang/go/blob/4b652e9f5f5c0793f2e41cd2876bce5a241b2c95/src/encoding/json/fuzz_test.go#L13

テストのテクニックの引き出しを増やしてイカしたテストを行ってください。

## 参考

https://go.dev/doc/security/fuzz/
