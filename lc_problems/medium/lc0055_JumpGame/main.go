func canJump(nums []int) bool {
	n := len(nums)
	if n == 1 {
		return true
	}
	maxIndex := nums[0]
	for i := 0; i <= maxIndex; i++ {
		newIndex := i + nums[i]
		if newIndex >= n-1 {
			return true
		}
		if newIndex > maxIndex {
			maxIndex = newIndex
		}
	}
	return false
}