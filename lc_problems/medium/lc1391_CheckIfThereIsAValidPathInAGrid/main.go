func hasValidPath(grid [][]int) bool {

	nRow := len(grid)
	nCol := len(grid[0])
	visited := make([][]bool, nRow)
	for i := 0; i < nRow; i++ {
		visited[i] = make([]bool, nCol)
	}

	isValid := func(row, col int) bool {
		if row < 0 || row >= nRow {
			return false
		}
		if col < 0 || col >= nCol {
			return false
		}
		if visited[row][col] {
			return false
		}
		return true
	}

	var helper func(row, col int) bool
	helper = func(row, col int) bool {
		if row == nRow-1 && col == nCol-1 {
			return true
		}

		visited[row][col] = true
		street := grid[row][col]
		if street == 1 {
			ret := false
			if isValid(row, col-1) && (grid[row][col-1] == 1 || grid[row][col-1] == 4 || grid[row][col-1] == 6) {
				ret = ret || helper(row, col-1)
			}
			if isValid(row, col+1) && (grid[row][col+1] == 1 || grid[row][col+1] == 3 || grid[row][col+1] == 5) {
				ret = ret || helper(row, col+1)
			}
			return ret
		}
		if street == 2 {
			ret := false
			if isValid(row-1, col) && (grid[row-1][col] == 2 || grid[row-1][col] == 3 || grid[row-1][col] == 4) {
				ret = ret || helper(row-1, col)
			}
			if isValid(row+1, col) && (grid[row+1][col] == 2 || grid[row+1][col] == 5 || grid[row+1][col] == 6) {
				ret = ret || helper(row+1, col)
			}
			return ret
		}
		if street == 3 {
			ret := false
			if isValid(row, col-1) && (grid[row][col-1] == 1 || grid[row][col-1] == 4 || grid[row][col-1] == 6) {
				ret = ret || helper(row, col-1)
			}
			if isValid(row+1, col) && (grid[row+1][col] == 2 || grid[row+1][col] == 5 || grid[row+1][col] == 6) {
				ret = ret || helper(row+1, col)
			}
			return ret
		}
		if street == 4 {
			ret := false
			if isValid(row, col+1) && (grid[row][col+1] == 1 || grid[row][col+1] == 3 || grid[row][col+1] == 5) {
				ret = ret || helper(row, col+1)
			}
			if isValid(row+1, col) && (grid[row+1][col] == 2 || grid[row+1][col] == 5 || grid[row+1][col] == 6) {
				ret = ret || helper(row+1, col)
			}
			return ret
		}
		if street == 5 {
			ret := false
			if isValid(row, col-1) && (grid[row][col-1] == 1 || grid[row][col-1] == 4 || grid[row][col-1] == 6) {
				ret = ret || helper(row, col-1)
			}
			if isValid(row-1, col) && (grid[row-1][col] == 2 || grid[row-1][col] == 3 || grid[row-1][col] == 4) {
				ret = ret || helper(row-1, col)
			}
			return ret
		}
		ret := false
		if isValid(row, col+1) && (grid[row][col+1] == 1 || grid[row][col+1] == 3 || grid[row][col+1] == 5) {
			ret = ret || helper(row, col+1)
		}
		if isValid(row-1, col) && (grid[row-1][col] == 2 || grid[row-1][col] == 3 || grid[row-1][col] == 4) {
			ret = ret || helper(row-1, col)
		}
		return ret
	}
	return helper(0, 0)
}