package main

import "fmt"

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}

func main() {
	arr := []int{1, 3, 5, 6}
	target1 := 2
	target2 := 5
	result1 := searchInsert(arr, target1)
	result2 := searchInsert(arr, target2)

	fmt.Println(result1, result2)
}
