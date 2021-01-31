func maximumWealth(accounts [][]int) int {
	result := 0

	for _, account := range accounts {
		sum := 0
		for _, val := range account {
			sum += val
		}
		if result < sum {
			result = sum
		}
	}
	return result
}