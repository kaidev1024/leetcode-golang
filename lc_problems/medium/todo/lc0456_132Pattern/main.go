func find132pattern(nums []int) bool {
	n := len(nums)
	second := -1 << 31

	stack := make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		if nums[i] < second {
			return true
		}
		ls := len(stack)
		for ls > 0 && nums[i] > stack[ls-1] {
			second = stack[ls-1]
			stack = stack[:ls-1]
			ls--
		}
		stack = append(stack, nums[i])
	}
	return false
}