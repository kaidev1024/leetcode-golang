func findBall(grid [][]int) []int {
	nr := len(grid)
	nc := len(grid[0])

	helper := func(col int) int {
		for i := 0; i < nr; i++ {
			if grid[i][col] == 1 {
				if col == nc-1 || grid[i][col+1] == -1 {
					return -1
				}
				col++
			} else {
				if col == 0 || grid[i][col-1] == 1 {
					return -1
				}
				col--
			}
		}
		return col
	}

	result := make([]int, nc)
	for i := 0; i < nc; i++ {
		result[i] = helper(i)
	}
	return result
}