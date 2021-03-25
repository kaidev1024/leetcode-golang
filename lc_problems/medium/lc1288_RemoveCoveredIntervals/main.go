func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		ii := intervals[i]
		ij := intervals[j]
		if ii[0] == ij[0] {
			return ii[1] > ij[1]
		}
		return ii[0] < ij[0]
	})

	n := len(intervals)
	cur := 0
	count := 1
	iter := 1
	for iter < n {
		if intervals[cur][1] >= intervals[iter][1] {
			iter++
			continue
		}
		count++
		cur = iter
		iter++
	}
	return count
}