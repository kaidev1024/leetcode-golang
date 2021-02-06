func oddCells(n int, m int, indices [][]int) int {
	rows := make([]int, n)
	cols := make([]int, m)
	l := len(indices)
	for i := 0; i < l; i++ {
		rows[indices[i][0]]++
		cols[indices[i][1]]++
	}

	result := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if (rows[i]+cols[j])%2 == 1 {
				result++
			}
		}
	}
	return result
}