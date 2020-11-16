func generate(numRows int) [][]int {
	result := make([][]int, numRows)

	for i := 1; i <= numRows; i++ {
		row := make([]int, i)
		row[0] = 1
		row[i-1] = 1
		for j := 1; j < i-1; j++ {
			row[j] = result[i-2][j-1] + result[i-2][j]
		}
		result[i-1] = row
	}

	return result
}