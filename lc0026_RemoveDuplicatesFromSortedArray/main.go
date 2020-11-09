package main

import "fmt"

func removeDuplicates(nums []int) int {
	l := len(nums)
	if l <= 1 {
		return l
	}

	slow, fast := 0, 1

	for fast < l {
		if nums[fast] == nums[fast-1] {
			fast++
		} else {
			slow++
			nums[slow] = nums[fast]
			fast++
		}
	}
	return slow + 1
}

func main() {
	nums := []int{1, 1, 2}
	size := removeDuplicates(nums)

	fmt.Println(size)
}
