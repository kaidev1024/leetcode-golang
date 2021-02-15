func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	n := len(height)
	l, r := 0, n-1
	result := 0
	for l < r {
		area := (r - l) * min(height[l], height[r])
		result = max(area, result)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return result
}