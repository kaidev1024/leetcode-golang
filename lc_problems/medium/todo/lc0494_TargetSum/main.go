func findTargetSumWays(nums []int, S int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum < S || (S+sum)%2 == 1 {
		return 0
	}
	target := (S + sum) >> 1
	n := len(nums)

	dp := make([][]int, target+1)
	for i := 0; i <= target; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= n; i++ {
		dp[0][i] = 1
	}

	for j := 1; j <= n; j++ {
		for i := 0; i <= target; i++ {
			dp[i][j] = dp[i][j-1]
			if i-nums[j-1] >= 0 {
				dp[i][j] += dp[i-nums[j-1]][j-1]
			}
		}
	}
	return dp[target][n]
}