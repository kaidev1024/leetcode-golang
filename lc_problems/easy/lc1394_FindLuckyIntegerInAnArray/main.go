func findLucky(arr []int) int {
	counts := make(map[int]int)

	for _, val := range arr {
		counts[val]++
	}

	result := -1
	for key, value := range counts {
		if key == value && key > result {
			result = key
		}
	}
	return result
}