func minSwaps(grid [][]int) int {
	nRows := len(grid)
	nCols := len(grid[0])

	lastOnes := make([]int, nRows)
	for row := 0; row < nRows; row++ {
		for col := nCols - 1; col >= 0; col-- {
			if grid[row][col] == 1 {
				lastOnes[row] = col
				break
			}
			if col == 0 {
				lastOnes[row] = -1
			}
		}
	}
	result := 0
	for row := 0; row < nRows; row++ {
		for i := row; i < len(lastOnes); i++ {
			if lastOnes[i] <= row {
				result += i - row
				for j := i; j > 0; j-- {
					lastOnes[j] = lastOnes[j-1]
				}
				break
			}
			if i == nRows-1 {
				return -1
			}
		}
	}
	return result
}
