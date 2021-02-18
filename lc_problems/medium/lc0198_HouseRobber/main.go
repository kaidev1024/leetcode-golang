func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}

	prev := nums[0]
	cur := max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		temp := max(nums[i]+prev, cur)
		prev = cur
		cur = temp
	}
	return cur
}