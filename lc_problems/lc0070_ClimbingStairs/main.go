package main

import "fmt"

func climbStairs(n int) int {
	if n < 2 {
		return n
	}

	first := 1
	second := 2
	var temp int
	for i := 3; i <= n; i++ {
		temp = second
		second += first
		first = temp
	}
	return second
}

func main() {
	n := 2
	result := climbStairs(n)
	fmt.Println(result)
}
