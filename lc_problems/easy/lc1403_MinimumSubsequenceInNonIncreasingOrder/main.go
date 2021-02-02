func minSubsequence(nums []int) []int {
	sort.Ints(nums)
	sum := 0
	for _, v := range nums {
		sum += v
	}

	l := len(nums)
	sum >>= 1

	s := 0
	result := make([]int, 0, l)
	for i := l - 1; i >= 0; i-- {
		s += nums[i]
		result = append(result, nums[i])
		if s > sum {
			break
		}
	}
	return result
}