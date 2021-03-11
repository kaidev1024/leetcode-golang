func minOperations(nums []int) int {
	helper := func(num int) (ones, twos int) {
		for num > 0 {
			if num&1 == 1 {
				ones++
				num--
			} else {
				twos++
				num = num >> 1
			}
		}
		return
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	totalOnes, maxTwos := 0, 0
	for _, num := range nums {
		ones, twos := helper(num)
		totalOnes += ones
		maxTwos = max(maxTwos, twos)
	}
	return totalOnes + maxTwos
}