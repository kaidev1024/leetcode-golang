func maxNonOverlapping(nums []int, target int) int {
	m := map[int]int{0: -1}

	var currentSum, count int
	previousIndex := -1

	// Prefix sum array
	for i := range nums {
		currentSum += nums[i]
		if index, ok := m[currentSum-target]; ok {
			if previousIndex <= index {
				count++
				previousIndex = i
			}
		}
		m[currentSum] = i
	}

	return count
}