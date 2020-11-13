package main

import "fmt"

func maxSubArray(nums []int) int {
	result := nums[0]

	localMax := 0
	for _, val := range nums {
		localMax += val
		if localMax > result {
			result = localMax
		}
		if localMax < 0 {
			localMax = 0
		}
	}
	return result
}

func main() {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	result := maxSubArray(arr)
	fmt.Println(result)
}
