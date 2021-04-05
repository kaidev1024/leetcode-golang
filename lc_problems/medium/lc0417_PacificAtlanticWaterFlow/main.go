func makeMatrix(m, n int) [][]bool {
	ret := make([][]bool, m)
	for i := 0; i < m; i++ {
		ret[i] = make([]bool, n)
	}
	return ret
}

func pacificAtlantic(heights [][]int) [][]int {
	nRow := len(heights)
	nCol := len(heights[0])

	P := makeMatrix(nRow, nCol)
	A := makeMatrix(nRow, nCol)
	dir := [][]int{[]int{0, 1}, []int{0, -1}, []int{1, 0}, []int{-1, 0}}
	var dfs func(matrix [][]int, start []int, result [][]bool)
	dfs = func(matrix [][]int, start []int, result [][]bool) {
		r, c := start[0], start[1]
		result[r][c] = true
		for i := 0; i < 4; i++ {
			ri := r + dir[i][0]
			ci := c + dir[i][1]
			if ri < 0 || ri >= nRow {
				continue
			}
			if ci < 0 || ci >= nCol {
				continue
			}
			if result[ri][ci] {
				continue
			}
			if matrix[ri][ci] < matrix[r][c] {
				continue
			}
			dfs(matrix, []int{ri, ci}, result)
		}
	}
	for i := 0; i < nCol; i++ {
		dfs(heights, []int{0, i}, P)
		dfs(heights, []int{nRow - 1, i}, A)
	}
	for i := 0; i < nRow-1; i++ {
		dfs(heights, []int{i + 1, 0}, P)
		dfs(heights, []int{i, nCol - 1}, A)
	}

	result := make([][]int, 0, nRow*nCol)
	for i := 0; i < nRow; i++ {
		for j := 0; j < nCol; j++ {
			if P[i][j] && A[i][j] {
				result = append(result, []int{i, j})
			}
		}
	}
	return result
}