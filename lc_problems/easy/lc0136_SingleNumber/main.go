package leetcode

func singleNumber(nums []int) int {
	s := 0
	for _, i := range nums {
		s ^= i
	}
	return s
}
