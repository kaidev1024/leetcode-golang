func numSpecial(mat [][]int) int {
	nRows := len(mat)
	nCols := len(mat[0])

	rows := make([]int, nRows)
	cols := make([]int, nCols)

	for i := 0; i < nRows; i++ {
		for j := 0; j < nCols; j++ {
			if mat[i][j] == 1 {
				if rows[i] > 0 {
					rows[i] = -rows[i]
				} else if rows[i] == 0 {
					rows[i] = j + 1
				}
				if cols[j] > 0 {
					cols[j] = -cols[j]
				} else if cols[j] == 0 {
					cols[j] = i + 1
				}
			}
		}
	}

	n := 0
	for _, v := range rows {
		if v > 0 && cols[v-1] > 0 {
			n++
		}
	}

	return n
}