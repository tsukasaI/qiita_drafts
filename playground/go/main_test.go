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
	defer fmt.Println("defer: 関数の終了時に実行")
	t.Cleanup(func() {
		fmt.Println("Cleanup: テストのクリーンアップ処理1")
	})
	t.Cleanup(func() {
		fmt.Println("Cleanup: テストのクリーンアップ処理2")
	})

	helperCleanup(t)

	// deferの登録

	fmt.Println("テスト処理の実行")

	t.Run("subtest", func(t *testing.T) {
		defer fmt.Println("defer: サブテストの終了時に実行")
		t.Cleanup(func() {
			fmt.Println("Cleanup: サブテストのクリーンアップ処理")
		})
		fmt.Println("サブテストの実行")
	})
	/*
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
	*/
}

func helperCleanup(t *testing.T) {
	t.Helper()
	t.Cleanup(func() {
		fmt.Println("Cleanup: ヘルパー関数のクリーンアップ処理")
	})
}
