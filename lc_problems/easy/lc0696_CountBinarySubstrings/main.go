func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func countBinarySubstrings(s string) int {
	result := 0
	n := len(s)
	if n < 2 {
		return 0
	}
	first, second := 1, 0
	for i := 1; i < n; i++ {
		if s[i] != s[i-1] {
			result += min(first, second)
			second = first
			first = 1
			continue
		}
		first++
	}
	result += min(first, second)
	return result
}