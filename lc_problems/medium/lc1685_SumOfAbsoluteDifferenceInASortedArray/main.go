func getSumAbsoluteDifferences(nums []int) []int {
	n := len(nums)
	sums := make([]int, n)
	sums[0] = nums[0]
	for i := 1; i < n; i++ {
		sums[i] = nums[i] + sums[i-1]
	}

	result := make([]int, n)

	result[0] = sums[n-1] - sums[0] - (n-1)*nums[0]
	result[n-1] = nums[n-1]*(n-1) - sums[n-2]

	for i := 1; i <= n-2; i++ {
		result[i] = (i*nums[i] - sums[i-1]) + (sums[n-1] - sums[i]) - nums[i]*(n-1-i)
	}
	return result
}