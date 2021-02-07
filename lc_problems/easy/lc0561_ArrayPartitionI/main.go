func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	s := 0
	for i, n := 0, len(nums); i < n; i += 2 {
		s += nums[i]
	}
	return s
}