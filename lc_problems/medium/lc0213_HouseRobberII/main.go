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
	for i := 2; i < n-1; i++ {
		temp := max(cur, prev+nums[i])
		prev = cur
		cur = temp
	}
	result := cur

	prev = nums[1]
	cur = max(nums[1], nums[2])
	for i := 3; i < n; i++ {
		temp := max(cur, prev+nums[i])
		prev = cur
		cur = temp
	}
	return max(result, cur)
}