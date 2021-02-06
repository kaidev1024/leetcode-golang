func checkStraightLine(coordinates [][]int) bool {
	n := len(coordinates)
	if n == 2 {
		return true
	}
	p0 := coordinates[0]
	p1 := coordinates[1]
	bx := p1[0] - p0[0]
	by := p1[1] - p0[1]
	helper := func(p []int) bool {
		dx := p[0] - p0[0]
		dy := p[1] - p0[1]
		if bx == 0 && dx == 0 {
			return true
		}
		if bx == 0 || dx == 0 {
			return false
		}
		if by == 0 && dy == 0 {
			return true
		}
		if by == 0 || dy == 0 {
			return false
		}
		return bx*dy == by*dx
	}

	for i := 2; i < n; i++ {
		if !helper(coordinates[i]) {
			return false
		}
	}
	return true
}