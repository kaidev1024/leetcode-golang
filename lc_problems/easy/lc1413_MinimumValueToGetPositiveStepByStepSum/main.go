func minStartValue(nums []int) int {
	minVal := 1
	sum := 0

	for _, val := range nums {
		sum += val
		if sum < minVal {
			minVal = sum
		}
	}

	if minVal >= 1 {
		return 1
	}
	return 1 - minVal
}