func numSubmat(mat [][]int) int {
	for i := 0; i < len(mat); i++ {
		for j := 1; j < len(mat[i]); j++ {
			if mat[i][j] == 1 {
				mat[i][j] += mat[i][j-1]
			}
		}
	}

	var r int

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 0 {
				break
			}
			r += mat[i][j]

			localMin := mat[i][j]

			for k := i + 1; k < len(mat); k++ {
				if mat[k][j] == 0 {
					break
				}

				localMin = min(localMin, mat[k][j])
				r += localMin
			}
		}
	}
	return r
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}