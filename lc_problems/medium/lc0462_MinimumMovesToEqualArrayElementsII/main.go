func minMoves2(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	if n <= 1 {
		return 0
	}
	a := nums[n>>1]
	result := 0
	dif := func(a, b int) int {
		if a > b {
			return a - b
		}
		return b - a
	}
	for _, v := range nums {
		result += dif(a, v)
	}
	return result
}