func maxRotateFunction(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}

	sum := 0
	r := 0
	for i, v := range nums {
		sum += v
		r += i * v
	}
	result := r
	for i := n - 1; i >= 0; i-- {
		r += sum - nums[i]*n
		result = max(result, r)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}