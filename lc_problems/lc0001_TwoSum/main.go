package main

import (
	"fmt"

	deepequal "github.com/kaidev1024/go-deep-equal"
)

func twoSum(nums []int, target int) []int {
	visited := make(map[int]int)
	for ind, val := range nums {
		pairVal := target - val
		if pairInd, ok := visited[pairVal]; ok {
			return []int{ind, pairInd}
		}
		visited[val] = ind
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	result := twoSum(nums, target)
	fmt.Println(deepequal.IsSliceEqual(result, []int{1, 0}))
}
