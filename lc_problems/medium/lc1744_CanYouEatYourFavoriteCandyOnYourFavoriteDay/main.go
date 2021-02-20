func canEat(candiesCount []int, queries [][]int) []bool {
	n := len(candiesCount)
	for i := 1; i < n; i++ {
		candiesCount[i] += candiesCount[i-1]
	}

	n = len(queries)
	result := make([]bool, n)

	for i, q := range queries {
		target := 0
		if q[0] != 0 {
			target = candiesCount[q[0]-1]
		}

		day := q[1] + 1
		fmt.Println(target, day)
		if day*q[2] <= target {
			result[i] = false
			continue
		}

		if q[1] >= candiesCount[q[0]] {
			result[i] = false
			continue
		}
		result[i] = true
	}
	return result
}
