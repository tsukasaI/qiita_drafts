package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestMain(m *testing.M) {
// 	fmt.Println("before tests")
// 	m.Run()
// 	fmt.Println("after tests")
// }

// func TestDeferVsCleanup(t *testing.T) {
// 	defer fmt.Println("defer: 関数の終了時に実行")
// 	t.Cleanup(func() {
// 		fmt.Println("Cleanup: テストのクリーンアップ処理1")
// 	})
// 	t.Cleanup(func() {
// 		fmt.Println("Cleanup: テストのクリーンアップ処理2")
// 	})

// 	helperCleanup(t)

// 	// deferの登録

// 	fmt.Println("テスト処理の実行")

// 	t.Run("subtest", func(t *testing.T) {
// 		defer fmt.Println("defer: サブテストの終了時に実行")
// 		t.Cleanup(func() {
// 			fmt.Println("Cleanup: サブテストのクリーンアップ処理")
// 		})
// 		fmt.Println("サブテストの実行")
// 	})
// 	/*
// 		before tests
// 		=== RUN   TestDeferVsCleanup
// 		テスト処理の実行
// 		=== RUN   TestDeferVsCleanup/subtest
// 		サブテストの実行
// 		defer: サブテストの終了時に実行
// 		Cleanup: サブテストのクリーンアップ処理
// 		defer: 関数の終了時に実行
// 		Cleanup: ヘルパー関数のクリーンアップ処理
// 		Cleanup: テストのクリーンアップ処理
// 		--- PASS: TestDeferVsCleanup (0.00s)
// 		    --- PASS: TestDeferVsCleanup/subtest (0.00s)
// 		PASS
// 		after tests
// 	*/
// }

// func helperCleanup(t *testing.T) {
// 	t.Helper()
// 	t.Cleanup(func() {
// 		fmt.Println("Cleanup: ヘルパー関数のクリーンアップ処理")
// 	})
// }

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

func FuzzIsEuqalIntSlice(f *testing.F) {
	cases := []struct{ a, b string }{
		{
			a: "abcd",
			b: "abcd",
		},
		// {
		// 	a: "aa",
		// 	b: "aaa",
		// },
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
