package leetcode

import "math"

func thirdMax(nums []int) int {
	a, b, c := math.MinInt64, math.MinInt64, math.MinInt64
	for _, v := range nums {
		if v == a || v == b || v == c {
			continue
		}
		if a == math.MinInt64 || a < v {
			c = b
			b = a
			a = v
		} else if b == math.MinInt64 || b < v {
			c = b
			b = v
		} else if c == math.MinInt64 || c < v {
			c = v
		}
	}
	if c == math.MinInt64 {
		return a
	}
	return c
}
