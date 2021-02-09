func sortedSquares(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	if nums[0] >= 0 {
		for i := 0; i < n; i++ {
			result[i] = nums[i] * nums[i]
		}
		return result
	}
	if nums[n-1] <= 0 {
		for i := 0; i < n; i++ {
			result[n-1-i] = nums[i] * nums[i]
		}
		return result
	}

	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	l, r := 0, n-1
	pos := n - 1
	for l <= r {
		if abs(nums[l]) > abs(nums[r]) {
			result[pos] = nums[l] * nums[l]
			l++
		} else {
			result[pos] = nums[r] * nums[r]
			r--
		}
		pos--
	}
	return result
}