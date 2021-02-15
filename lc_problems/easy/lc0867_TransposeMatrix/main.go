func transpose(A [][]int) [][]int {
	B := make([][]int, len(A[0]))
	for i := range B {
		B[i] = make([]int, len(A))
		for j := range B[i] {
			B[i][j] = A[j][i]
		}
	}

	return B
}