func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxVal := 0
	n := len(candies)
	result := make([]bool, n)

	for i := 0; i < n; i++ {
		if candies[i] > maxVal {
			maxVal = candies[i]
		}
	}

	for i := 0; i < n; i++ {
		if candies[i]+extraCandies >= maxVal {
			result[i] = true
		}
	}
	return result
}