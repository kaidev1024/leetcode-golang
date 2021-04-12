func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}

	target := sum >> 1
	dp := make([]bool, target+1)
	dp[0] = true

	for _, num := range nums {
		for i := target; i >= num; i-- {
			if i >= num && dp[i-num] {
				dp[i] = true
			}
		}
	}
	return dp[target]
}