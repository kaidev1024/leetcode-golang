func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	if n < 2 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		ii := intervals[i]
		ij := intervals[j]
		if ii[0] == ij[0] {
			return ii[1] < ij[1]
		}
		return ii[0] < ij[0]
	})

	cur := intervals[0]
	result := 0
	for i := 1; i < n; i++ {
		if intervals[i][0] >= cur[1] {
			cur = intervals[i]
			continue
		}
		result++
		if intervals[i][1] < cur[1] {
			cur = intervals[i]
		}

	}
	return result
}