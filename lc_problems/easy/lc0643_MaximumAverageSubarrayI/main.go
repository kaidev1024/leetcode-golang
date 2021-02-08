func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	maxVal := sum

	l := len(nums)
	for i := k; i < l; i++ {
		sum += nums[i] - nums[i-k]
		if sum > maxVal {
			maxVal = sum
		}
	}
	return float64(maxVal) / float64(k)
}