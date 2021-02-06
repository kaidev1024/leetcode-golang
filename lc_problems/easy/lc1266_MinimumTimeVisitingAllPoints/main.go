func minTimeToVisitAllPoints(points [][]int) int {
	l := len(points)
	if l == 1 {
		return 0
	}

	helper := func(x, y int) int {
		if x < 0 {
			x = -x
		}
		if y < 0 {
			y = -y
		}
		if x > y {
			return x
		}
		return y
	}

	r := 0

	for i := 1; i < l; i++ {
		dx := points[i][0] - points[i-1][0]
		dy := points[i][1] - points[i-1][1]
		r += helper(dx, dy)
	}

	return r
}