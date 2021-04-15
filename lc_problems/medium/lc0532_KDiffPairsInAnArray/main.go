func findPairs(nums []int, k int) int {
	result := 0
	if k == 0 {
		counts := make(map[int]int)
		for _, v := range nums {
			counts[v]++
		}
		for _, v := range counts {
			if v > 1 {
				result++
			}
		}
		return result
	}
	visited := make(map[int]struct{})

	for _, v := range nums {
		if _, ok := visited[v]; ok {
			continue
		}
		visited[v] = struct{}{}
		if _, ok := visited[v+k]; ok {
			result++
		}
		if _, ok := visited[v-k]; ok {
			result++
		}
	}
	return result
}