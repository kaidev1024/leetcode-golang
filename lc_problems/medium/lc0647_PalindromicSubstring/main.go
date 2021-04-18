func countSubstrings(s string) int {
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	for i := 0; i < n; i++ {
		dp[i][i] = true
	}
	ret := n

	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			ret++
			dp[i][i+1] = true
		}
	}
	for l := 3; l <= n; l++ {
		for start := 0; start < n; start++ {
			end := start + l - 1
			if end == n {
				break
			}
			if s[start] != s[end] {
				continue
			}
			if dp[start+1][end-1] {
				dp[start][end] = true
				ret++
			}
		}
	}
	return ret
}