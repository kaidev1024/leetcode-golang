func diagonalSort(mat [][]int) [][]int {

	m := len(mat)
	n := len(mat[0])

	diagonalsMap := make(map[int][]int)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			diagonalsMap[i-j] = append(diagonalsMap[i-j], mat[i][j])
		}
	}

	for _, v := range diagonalsMap {
		sort.Ints(v)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			mat[i][j] = diagonalsMap[i-j][0]
			diagonalsMap[i-j] = diagonalsMap[i-j][1:]
		}
	}

	return mat
}