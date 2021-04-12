func findDiagonalOrder(mat [][]int) []int {
	nRow := len(mat)
	nCol := len(mat[0])
	n := nRow * nCol

	ret := make([]int, 0, n)
	k := nRow + nCol - 1
	up := true
	for i := 0; i < k; i++ {
		if up {
			r, c := i, 0
			if i >= nRow {
				r = nRow - 1
				c = i - r
			}
			for r >= 0 && c < nCol {
				ret = append(ret, mat[r][c])
				r--
				c++
			}
			up = false
		} else {
			r, c := 0, i
			if i >= nCol {
				c = nCol - 1
				r = i - c
			}
			for r < nRow && c >= 0 {
				ret = append(ret, mat[r][c])
				r++
				c--
			}
			up = true
		}
	}
	return ret
}