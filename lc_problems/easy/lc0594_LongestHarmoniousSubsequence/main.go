func findLHS(nums []int) int {
	counts := make(map[int]int)
	for _, v := range nums {
		counts[v]++
	}

	maxVal := 0
	for k, v := range counts {
		if count, ok := counts[k+1]; ok {
			result := count + v
			if result > maxVal {
				maxVal = result
			}
		}
	}
	return maxVal
}