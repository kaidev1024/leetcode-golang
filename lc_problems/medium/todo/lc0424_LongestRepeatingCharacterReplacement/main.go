func characterReplacement(s string, k int) int {
	n := len(s)
	start := 0
	maxCount := 0
	counts := make([]int, 26)
	result := 0
	for end := 0; end < n; end++ {
		pos := s[end] - 'A'
		counts[pos]++
		maxCount = max(maxCount, counts[pos])
		for start <= end && end-start+1-maxCount > k {
			counts[s[start]-'A']--
			start++
		}
		result = max(result, end-start+1)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}