func productExceptSelf(nums []int) []int {
	n := len(nums)
	l := make([]int, n)
	r := make([]int, n)
	l[0] = 1
	r[n-1] = 1
	for i := 0; i < n-1; i++ {
		l[i+1] = l[i] * nums[i]
		r[n-2-i] = r[n-1-i] * nums[n-1-i]
	}

	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = l[i] * r[i]
	}
	return result
}