func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}

	result := []int{1, 1}
	for i := 1; i < rowIndex; i++ {
		nextRow := make([]int, i+2)
		nextRow[0] = 1
		nextRow[i+1] = 1
		for j := 1; j <= i; j++ {
			nextRow[j] = result[j] + result[j-1]
		}
		result = nextRow
	}
	return result
}