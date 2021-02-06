func decompressRLElist(nums []int) []int {
	l := len(nums)
	n := 0
	for i := 0; i < l; i += 2 {
		n += nums[i]
	}

	result := make([]int, n)
	pos := 0
	for i := 0; i < l; i += 2 {
		for j := 0; j < nums[i]; j++ {
			result[pos] = nums[i+1]
			pos++
		}
	}
	return result
}