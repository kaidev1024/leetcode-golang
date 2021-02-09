func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func surfaceArea(grid [][]int) int {
	sum := 0
	count := 0
	n := len(grid)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 0 {
				sum += grid[i][j]*4 + 2
			}
			if i != 0 {
				count += min(grid[i-1][j], grid[i][j])
			}
			if j != 0 {
				count += min(grid[i][j-1], grid[i][j])
			}
		}
	}

	return sum - count*2
}