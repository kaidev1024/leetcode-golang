func countVowelStrings(n int) int {
	result1 := make([]int, 5)
	result2 := make([]int, 5)

	for i := 0; i < 5; i++ {
		result1[i] = 1
	}

	for i := 1; i <= n; i++ {
		result2[0] = 1
		for j := 1; j < 5; j++ {
			result2[j] = result2[j-1] + result1[j]
		}
		result1, result2 = result2, result1
	}

	return result1[4]
}