package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(adder(1, 2))
	fmt.Println(adder(1, 2))
	fmt.Println(adder(1, 2))
	fmt.Println(adder(1, 2))
}

func adder(a, b int) int {
	time.Sleep(100 * time.Millisecond)
	return a + b
}
