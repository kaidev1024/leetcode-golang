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

func maxProduct(nums []int) int {
	minV, maxV := 1, 1
	r := nums[0]
	for _, v := range nums {
		if v < 0 {
			minV, maxV = maxV, minV
		}
		minV = min(minV*v, v)
		maxV = max(maxV*v, v)
		r = max(r, maxV)
	}
	return r
}