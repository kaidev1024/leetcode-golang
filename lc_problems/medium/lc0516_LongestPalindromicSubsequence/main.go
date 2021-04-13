func longestPalindromeSubseq(s string) int {
	dp := make([][]int, len(s))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(s))
	}

	for i := 0; i < len(dp); i++ {
		dp[i][i] = 1
	}

	for i := 0; i < len(dp)-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = 2
		} else {
			dp[i][i+1] = 1
		}
	}

	for i := 2; i < len(s); i++ {
		for j := 0; j+i < len(s); j++ {
			if s[j] != s[j+i] {
				dp[j][j+i] = max(dp[j+1][j+i], dp[j][j+i-1])
			} else {
				dp[j][j+i] = 2 + dp[j+1][j+i-1]
			}
		}
	}
	return dp[0][len(s)-1]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}