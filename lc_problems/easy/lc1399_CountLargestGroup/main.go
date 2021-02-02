func countLargestGroup(n int) int {
	sum := func(v int) int {
		result := 0
		for v > 0 {
			result += v % 10
			v /= 10
		}
		return result
	}

	counts := make([]int, 37)
	for i := 1; i <= n; i++ {
		counts[sum(i)]++
	}

	maxVal := 0
	count := 0

	for _, v := range counts {
		if v > maxVal {
			maxVal = v
			count = 1
			continue
		}
		if v == maxVal {
			count++
		}
	}

	return count
}