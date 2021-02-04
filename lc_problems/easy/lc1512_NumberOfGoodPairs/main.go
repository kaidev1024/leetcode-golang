func numIdenticalPairs(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	result := 0
	for _, v := range m {
		result += v * (v - 1) >> 1
	}

	return result
}