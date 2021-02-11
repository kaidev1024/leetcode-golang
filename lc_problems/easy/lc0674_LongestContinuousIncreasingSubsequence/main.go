func findLengthOfLCIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	l := 1
	maxL := 0

	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			l++
		} else {
			if l > maxL {
				maxL = l
			}
			l = 1
		}
	}
	if l > maxL {
		maxL = l
	}

	return maxL

}