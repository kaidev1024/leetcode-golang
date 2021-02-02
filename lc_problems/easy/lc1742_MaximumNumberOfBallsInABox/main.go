func countBalls(lowLimit int, highLimit int) int {
	sum := func(n int) int {
		result := 0
		for n > 0 {
			result += n % 10
			n /= 10
		}
		return result
	}

	counts := make([]int, 46)

	maxVal := 0
	for i := lowLimit; i <= highLimit; i++ {
		s := sum(i)
		counts[s]++
		if counts[s] > maxVal {
			maxVal = counts[s]
		}
	}

	return maxVal
}