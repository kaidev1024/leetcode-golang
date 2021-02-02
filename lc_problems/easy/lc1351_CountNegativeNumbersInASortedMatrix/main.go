func countNegatives(grid [][]int) int {
	nRows := len(grid)
	if nRows == 0 {
		return 0
	}

	nCols := len(grid[0])

	result := 0

	j := nCols - 1
	for i := 0; i < nRows; i++ {
		for j >= 0 {

			if grid[i][j] < 0 {
				j--
			} else {

				result += nCols - j - 1
				break
			}
		}
		if j == -1 {
			result += nCols
		}
	}
	return result
}