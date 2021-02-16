func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func insert(intervals [][]int, newInterval []int) [][]int {
	n := len(intervals)
	if n == 0 {
		return [][]int{newInterval}
	}
	l, r := 0, n-1

	start := 0
	firstHalf := [][]int{newInterval}
	if newInterval[0] >= intervals[0][0] {
		for l < r {
			mid := (l + r + 1) >> 1
			if intervals[mid][0] > newInterval[0] {
				r = mid - 1
			} else {
				l = mid
			}
		}
		firstHalf = make([][]int, l+1)
		copy(firstHalf, intervals)
		if firstHalf[l][1] >= newInterval[0] {
			firstHalf[l][1] = max(firstHalf[l][1], newInterval[1])
			start = l + 1
		} else {
			firstHalf = append(firstHalf, newInterval)
			start = l + 1
			l++
		}

	}

	for start < n {
		if firstHalf[l][1] >= intervals[start][0] {
			firstHalf[l][1] = max(firstHalf[l][1], intervals[start][1])
			start++
		} else {
			break
		}
	}
	firstHalf = append(firstHalf, intervals[start:]...)
	return firstHalf
}