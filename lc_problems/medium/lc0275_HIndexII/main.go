func hIndex(citations []int) int {
	n := len(citations)
	if n == 0 {
		return 0
	}
	l, r := 0, n-1

	for l < r {
		mid := (l + r) / 2
		cite := citations[mid]
		count := n - mid
		if cite >= count {
			r = mid
		} else {
			l = mid + 1
		}
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	return min(n-l, citations[l])
}