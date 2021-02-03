func shuffle(nums []int, n int) []int {
	l := len(nums)
	result := make([]int, 0, l)

	for i := 0; i < n; i++ {
		for j := i; j < l; j += n {
			result = append(result, nums[j])
		}
	}

	return result
}