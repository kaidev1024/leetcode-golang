func searchMatrix(matrix [][]int, target int) bool {
	n := len(matrix)
	m := len(matrix[0])
	r := 0
	c := m - 1

	for r < n && c >= 0 {
		if matrix[r][c] == target {
			return true
		}
		if matrix[r][c] > target {
			c--
			continue
		}
		r++
	}
	return false
}