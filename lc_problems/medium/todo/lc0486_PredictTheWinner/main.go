func PredictTheWinner(nums []int) bool {
	return helper(nums, 0, 1)
}

func helper(nums []int, diff int, player int) bool {
	n := len(nums)
	if n == 0 {
		return diff > 0 || (diff == 0 && player == 1)
	}

	if helper(nums[:n-1], -diff-nums[n-1], 3-player) && helper(nums[1:], -diff-nums[0], 3-player) {
		return false
	}
	return true
}