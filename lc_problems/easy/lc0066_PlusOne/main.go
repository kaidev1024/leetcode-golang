package main

import "fmt"

func plusOne(digits []int) []int {
	l := len(digits)
	for i := l - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] += 1
			return digits
		} else {
			digits[i] = 0
		}
	}

	result := make([]int, l+1)
	result[0] = 1
	return result
}

func main() {
	arr := []int{1, 2, 3}
	result := plusOne(arr)
	fmt.Println(result)
}
