func findRightInterval(intervals [][]int) []int {
	n := len(intervals)
	if n == 1 {
		return []int{-1}
	}
	m := make(map[int]int)
	starts := make([]int, n)
	for i, v := range intervals {
		starts[i] = v[0]
		m[v[0]] = i
	}

	sort.Ints(starts)

	result := make([]int, n)
	for i := 0; i < n; i++ {
		index := sort.SearchInts(starts, intervals[i][1])
		if index == n {
			result[i] = -1
		} else {
			result[i] = m[starts[index]]
		}
	}

	return result
}