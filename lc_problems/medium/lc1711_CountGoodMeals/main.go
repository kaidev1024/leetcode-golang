func countPairs(deliciousness []int) int {
	counts := make(map[int]int)

	for _, v := range deliciousness {
		counts[v]++
	}

	arr := make([]int, 22)
	for i := 0; i <= 21; i++ {
		arr[i] = 1 << i
	}

	result := 0

	for key, value := range counts {
		for i := 0; i < 22; i++ {
			if key > arr[i] {
				continue
			}
			target := arr[i] - key
			if count, ok := counts[target]; ok {
				if target == key {
					result += (value - 1) * value / 2
				} else {
					result += value * count
				}
				result %= 1e9 + 7
			}
		}
		delete(counts, key)
	}
	return result
}