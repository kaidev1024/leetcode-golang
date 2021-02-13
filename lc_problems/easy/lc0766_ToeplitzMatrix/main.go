func isToeplitzMatrix(matrix [][]int) bool {
	nr := len(matrix)
	nc := len(matrix[0])

	for i := 1; i < nr; i++ {
		for j := 1; j < nc; j++ {
			if matrix[i][j] != matrix[i-1][j-1] {
				return false
			}
		}
	}
	return true
}