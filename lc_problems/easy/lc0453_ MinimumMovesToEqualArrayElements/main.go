package leetcode

func minMoves(nums []int) int {
	l := len(nums)

	minVal := nums[0]
	sum := nums[0]
	for i := 1; i < l; i++ {
		cur := nums[i]
		if cur < minVal {
			minVal = cur
		}
		sum += cur
	}

	return sum - minVal*l
}
