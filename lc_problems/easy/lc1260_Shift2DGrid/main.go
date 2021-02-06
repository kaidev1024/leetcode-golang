func shiftGrid(grid [][]int, k int) [][]int {
	nr := len(grid)
	nc := len(grid[0])

	n := nr * nc

	k %= n

	rr := k / nc
	cc := k % nc

	result := make([][]int, nr)
	for i := 0; i < nr; i++ {
		result[i] = make([]int, nc)
	}

	for i, r, c := 0, 0, 0; i < n; i++ {
		result[rr][cc] = grid[r][c]
		cc++
		if cc == nc {
			cc = 0
			rr++
			if rr == nr {
				rr = 0
			}
		}

		c++
		if c == nc {
			c = 0
			r++
		}
	}
	return result
}