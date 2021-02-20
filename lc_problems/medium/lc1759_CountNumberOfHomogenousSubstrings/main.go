func helper(n int) int {
	return ((n + 1) * n >> 1) % 1000000007
}

func countHomogenous(s string) int {
	n := len(s)
	start := 0
	end := 0
	result := 0
	for end < n {
		if s[end] == s[start] {
			end++
		} else {
			result += helper(end - start)
			start = end
		}
	}
	result += helper(end - start)
	return result % 1000000007
}