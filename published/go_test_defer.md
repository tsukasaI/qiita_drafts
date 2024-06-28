# Goのテストの後処理におけるdeferとt.Cleanupの差

単体テストにおいて前処理と後処理をする場合が多い。

例えばDBのデータを作成/削除する、一時ファイルを作成/削除するなどテスト対象の関数やメソッドの責務のみに集中するためには、テストの前処理と後処理を行う必要がある。

Goのテストにおいての前処理と後処理にはTestMain関数を用いる。

一方で各テスト関数において前処理と後処理を行う方法として、`defer`と`t.Cleanup`がある。

この挙動の差をサンプルコードを用いて確認する。

## deferとt.Cleanupの違い

`defer`は関数の終了時に実行される処理を登録する。

一方で`t.Cleanup`はテスト関数の後処理を登録する関数で、テスト関数が終了するときに登録した関数が実行される。

## サンプルのテストコード

さて問題です。

次のようなテストコードが実行されるとどのように表示されるか。

```go
package main_test

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("before tests")
	m.Run()
	fmt.Println("after tests")
}

func TestDeferVsCleanup(t *testing.T) {
	t.Cleanup(func() {
		fmt.Println("Cleanup: テストのクリーンアップ処理")
	})

	helperCleanup(t)

	// deferの登録
	defer fmt.Println("defer: 関数の終了時に実行")

	fmt.Println("テスト処理の実行")

	t.Run("subtest", func(t *testing.T) {
		defer fmt.Println("defer: サブテストの終了時に実行")
		t.Cleanup(func() {
			fmt.Println("Cleanup: サブテストのクリーンアップ処理")
		})
		fmt.Println("サブテストの実行")
	})
}

func helperCleanup(t *testing.T) {
	t.Helper()
	t.Cleanup(func() {
		fmt.Println("Cleanup: ヘルパー関数のクリーンアップ処理")
	})
}
```

---

答え

```sh
$ go test -v
before tests
=== RUN   TestDeferVsCleanup
テスト処理の実行
=== RUN   TestDeferVsCleanup/subtest
サブテストの実行
defer: サブテストの終了時に実行
Cleanup: サブテストのクリーンアップ処理
defer: 関数の終了時に実行
Cleanup: ヘルパー関数のクリーンアップ処理
Cleanup: テストのクリーンアップ処理
--- PASS: TestDeferVsCleanup (0.00s)
    --- PASS: TestDeferVsCleanup/subtest (0.00s)
PASS
after tests
```

## 解説

Goのテストやベンチマーク測定が実行されると、最初にTestMain(m *testing.M)関数が実行される。

メインのgoroutineで実行され、前処理と後処理を行うことができる。

`m.Run()`を実行することで各テスト関数が実行されるため befer testsが最初に、after tests が最後に表示される。

TestDeferVsCleanup関数を順に見ていく。

t.Cleanup、及びt.CleanupをコールしているhelperCleanup はテスト関数の後処理を登録する関数で、テスト関数が終了するときに登録した関数が実行される。

関数の終了時に実行されるdeferは、関数の終了時に実行されるがt.Cleanupよりも先に実行される。

次にfmt.Println("テスト処理の実行")が書かれているのでこれは最初に表示される。

続いてサブテストが実行され、fmt.Println("サブテストの実行")が書かれているため次はこれが表示される。

サブテスト内にdeferとt.Cleanupが登録されているため、サブテストの後処理がdeferで登録されたfmt.Println("defer: サブテストの終了時に実行") -> t.Cleanupで登録されたfmt.Println("Cleanup: サブテストのクリーンアップ処理")の順に表示される。

サブテストの後処理が終わると、TestDeferVsCleanup関数の後処理が実行される。

TestDeferVsCleanup関数の後処理はdeferで登録されたfmt.Println("defer: 関数の終了時に実行") -> t.Cleanupで登録されたfmt.Println("Cleanup: ヘルパー関数のクリーンアップ処理")の順に表示される。

t.Clearupで登録された関数は、登録された順と逆順に実行される。

参考: https://pkg.go.dev/testing#hdr-Main

## まとめ

defer, t.Cleanupはテスト関数の後処理を書け、サブテスト内やヘルパー関数内でも登録した場合の動きを見た。

これらの性質を理解して、テストコードを書く際に使い分けていきましょう！
