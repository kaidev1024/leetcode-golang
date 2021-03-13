func getLastMoment(n int, left []int, right []int) int {
	maxLeft := 0
	minRight := n
	for _, v := range left {
		if v > maxLeft {
			maxLeft = v
		}
	}
	for _, v := range right {
		if v < minRight {
			minRight = v
		}
	}
	d := n - minRight
	if d > maxLeft {
		return d
	}
	return maxLeft
}