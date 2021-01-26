package leetcode

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	l := len(nums)
	start := 0
	end := 0

	result := make([]string, 0, l)
	for start < l {
		for end+1 < l && nums[end+1]-nums[end] == 1 {
			end++
		}

		var s string
		if start == end {
			s = strconv.Itoa(nums[start])
		} else {
			s = fmt.Sprintf("%v->%v", nums[start], nums[end])
		}

		result = append(result, s)
		start = end + 1
		end = start
	}
	return result
}
