func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for _, str := range strs {
		count0, count1 := count(str)
		for i := m; i >= count0; i-- {
			for j := n; j >= count1; j-- {
				dp[i][j] = max(dp[i][j], 1+dp[i-count0][j-count1])
			}
		}
	}

	return dp[m][n]
}

func count(s string) (int, int) {
	zeros, ones := 0, 0
	for _, v := range s {
		if v == '0' {
			zeros++
		} else {
			ones++
		}
	}
	return zeros, ones
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}