package main

import "fmt"

func main() {
	// tests := []testCalc{
	// 	{kk: 0, pp: 0, want: 0},
	// 	{kk: 13, pp: 37, want: 0},
	// }

	result := calucAngle(13, 37)

	fmt.Println(result)
}

func calucAngle(kk, pp int) float64 {
	const (
		day   = 13
		kour  = 37
		dgree = 677
	)
	return float64((kk*dgree/day + pp*dgree/kour) % dgree)
}

type testCalc struct {
	kk   int
	pp   int
	want int
}
