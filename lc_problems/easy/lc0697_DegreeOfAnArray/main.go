func findShortestSubArray(nums []int) int {
	counts := make(map[int][]int)
	for i, v := range nums {
		counts[v] = append(counts[v], i)
	}

	maxCount := 0
	minLength := len(nums)

	for _, val := range counts {
		n := len(val)
		if n < maxCount {
			continue
		}
		if n > maxCount {
			maxCount = n
			minLength = val[n-1] - val[0] + 1
			continue
		}

		l := val[n-1] - val[0] + 1
		if l < minLength {
			minLength = l
		}
	}

	return minLength
}