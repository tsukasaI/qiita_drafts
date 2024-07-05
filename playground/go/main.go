package main

import (
	"fmt"
)

func main() {
	ints := []int{1, 2, 3, 4, 5}
	fmt.Println("[]int{1, 2, 3, 4, 5}のlen:", len(ints))
	fmt.Println("[]int{1, 2, 3, 4, 5}のrange")
	for i, v := range ints {
		fmt.Println(i, v)
	}

	var nilSlice []int
	fmt.Println("nilSliceのlen:", len(nilSlice))
	fmt.Println("nilSliceのrange")
	for i, v := range nilSlice {
		fmt.Println(i, v)
	}

	emptySlice := []int{}
	fmt.Println("[]int{}のlen:", len(emptySlice))
	fmt.Println("[]int{}のrange")
	for i, v := range emptySlice {
		fmt.Println(i, v)
	}

	kv := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("map[string]int{\"a\": 1, \"b\": 2, \"c\": 3}のlen:", len(kv))
	fmt.Println("map[string]int{\"a\": 1, \"b\": 2, \"c\": 3}のrange")
	for k, v := range kv {
		fmt.Println(k, v)
	}

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
}
