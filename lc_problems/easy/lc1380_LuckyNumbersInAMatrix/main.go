func luckyNumbers(matrix [][]int) []int {
	nRow := len(matrix)
	if nRow == 0 {
		return nil
	}
	nCol := len(matrix[0])

	rowMin := make([]int, nRow)
	for i := 0; i < nRow; i++ {
		rowMin[i] = 100000
	}
	colMax := make([]int, nCol)
	for i := 0; i < nCol; i++ {
		colMax[i] = 1
	}

	for i := 0; i < nRow; i++ {
		for j := 0; j < nCol; j++ {
			if matrix[i][j] < rowMin[i] {
				rowMin[i] = matrix[i][j]
			}
			if matrix[i][j] > colMax[j] {
				colMax[j] = matrix[i][j]
			}
		}
	}

	result := make([]int, 0, nRow*nCol)
	for i := 0; i < nRow; i++ {
		for j := 0; j < nCol; j++ {
			if rowMin[i] == colMax[j] {
				result = append(result, rowMin[i])
			}
		}
	}
	return result
}