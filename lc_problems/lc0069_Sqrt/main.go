package main

import "fmt"

func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	l := 1
	r := x

	for l < r {
		mid := l + (r-l)/2
		div := x / mid
		if mid == div {
			return mid
		}

		if mid < div {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	if l > x/l {
		return l - 1
	}
	return l
}

func main() {
	x := 8
	result := mySqrt(x)
	fmt.Println(result)
}
