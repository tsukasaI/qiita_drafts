package main

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("Before running tests")

	m.Run()

	fmt.Println("After running tests")
}

func TestAdder(t *testing.T) {
	defer fmt.Println("defer")

	cases := map[string]struct {
		a, b, expected int
	}{
		"1": {1, 2, 3},
		"2": {2, 3, 5},
		"3": {3, 4, 7},
	}

	for name, c := range cases {
		t.Run(fmt.Sprintf("%d+%d=%d", c.a, c.b, c.expected), func(t *testing.T) {
			t.Parallel()
			got := adder(c.a, c.b)
			if got != c.expected {
				t.Errorf("adder(%d, %d) == %d, want %d", c.a, c.b, got, c.expected)
			}
			fmt.Println("Test case", name, "done")
		})
	}
}
