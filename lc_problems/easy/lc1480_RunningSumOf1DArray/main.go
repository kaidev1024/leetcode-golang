func runningSum(nums []int) []int {
	for i, l := 1, len(nums); i < l; i++ {
		nums[i] += nums[i-1]
	}
	return nums
}