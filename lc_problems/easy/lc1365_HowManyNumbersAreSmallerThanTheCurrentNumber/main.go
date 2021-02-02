func smallerNumbersThanCurrent(nums []int) []int {
	counts := make([]int, 101)
	n := len(nums)

	for i := 0; i < n; i++ {
		counts[nums[i]]++
	}

	n1 := n
	for i := 100; i >= 0; i-- {
		if counts[i] > 0 {
			counts[i] = n1 - counts[i]
			n1 = counts[i]
		}
	}

	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = counts[nums[i]]
	}
	return result
}