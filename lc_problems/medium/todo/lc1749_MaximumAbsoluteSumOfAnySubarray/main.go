func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxAbsoluteSum(nums []int) int {
	n := len(nums)
	gMin := nums[0]
	gMax := nums[0]

	sumMin := 0
	sumMax := 0

	for i := 0; i < n; i++ {
		sumMax += nums[i]
		gMax = max(sumMax, gMax)
		if sumMax < 0 {
			sumMax = 0
		}
		sumMin += nums[i]
		gMin = min(sumMin, gMin)
		if sumMin > 0 {
			sumMin = 0
		}
	}
	return max(abs(gMax), abs(gMin))
}