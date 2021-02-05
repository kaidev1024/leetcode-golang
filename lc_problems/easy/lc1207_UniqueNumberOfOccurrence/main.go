func uniqueOccurrences(arr []int) bool {
	m := make(map[int]int)
	for _, v := range arr {
		m[v]++
	}

	counts := make(map[int]struct{})
	for _, v := range m {
		if _, ok := counts[v]; ok {
			return false
		}
		counts[v] = struct{}{}
	}
	return true
}