package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	minVal := math.MinInt32
	maxVal := math.MaxInt32

	var result int
	for x != 0 {
		result *= 10
		result += x % 10
		if result < minVal || result > maxVal {
			return 0
		}
		x /= 10
	}

	return result
}

func main() {
	x := 2371
	fmt.Println(x, reverse(x))
}
