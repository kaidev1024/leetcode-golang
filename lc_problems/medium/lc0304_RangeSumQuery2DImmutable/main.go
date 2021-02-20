type NumMatrix struct {
	sums [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	nr := len(matrix)
	if nr == 0 {
		return NumMatrix{nil}
	}
	nc := len(matrix[0])
	sums := make([][]int, nr+1)
	for i := 0; i <= nr; i++ {
		sums[i] = make([]int, nc+1)
	}

	for i := 1; i <= nr; i++ {
		for j := 1; j <= nc; j++ {
			sums[i][j] = sums[i-1][j] + sums[i][j-1] - sums[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	return NumMatrix{sums}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if len(this.sums) == 0 {
		return 0
	}
	return this.sums[row2+1][col2+1] + this.sums[row1][col1] - this.sums[row2+1][col1] - this.sums[row1][col2+1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */