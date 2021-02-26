func minOperations(nums []int, x int) int {
	length, s := len(nums), 0
	for _, n := range nums {
		s += n
	}
	if s < x {
		return -1
	}
	if s == x {
		return length
	}
	target, maxLen := s-x, -1
	l, r, window := 0, 0, 0
	for r < length {
		window += nums[r]
		r++
		for window >= target {
			if window == target {
				maxLen = max(maxLen, r-l)
			}
			window -= nums[l]
			l++
		}
	}
	if maxLen == -1 {
		return -1
	}
	return length - maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}