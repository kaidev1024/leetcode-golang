func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	v := 1
	left, right := 0, n-1
	top, bottom := 0, n-1

	n *= n
	for v <= n {
		for i := left; i <= right; i++ {
			matrix[top][i] = v
			v++
		}
		top++
		for i := top; i <= bottom; i++ {
			matrix[i][right] = v
			v++
		}
		right--
		for i := right; i >= left; i-- {
			matrix[bottom][i] = v
			v++
		}
		bottom--
		for i := bottom; i >= top; i-- {
			matrix[i][left] = v
			v++
		}
		left++
	}
	return matrix
}