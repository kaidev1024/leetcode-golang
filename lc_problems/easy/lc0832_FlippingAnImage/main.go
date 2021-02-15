func flipAndInvertImage(A [][]int) [][]int {
	for i := 0; i < len(A); i++ {
		row := A[i]
		last := len(row) - 1
		for j := 0; j <= last/2; j++ {
			row[j], row[last-j] = row[last-j]^1, row[j]^1
		}
	}
	return A
}