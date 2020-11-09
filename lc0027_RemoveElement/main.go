package main

import "fmt"

func removeElement(nums []int, val int) int {
	l := len(nums)
	slow := 0
	fast := 0
	for fast < l {
		if nums[fast] == val {
			fast++
		} else {
			nums[slow] = nums[fast]
			slow++
			fast++
		}
	}
	return slow
}

func main() {
	arr := []int{3, 2, 2, 3}
	ele := 3

	result := removeElement(arr, ele)

	fmt.Println(result)
}
