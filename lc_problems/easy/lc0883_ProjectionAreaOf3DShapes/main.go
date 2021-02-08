func projectionArea(grid [][]int) int {
	n := len(grid)

	sum := 0

	arrx := make([]int, n)
	arry := make([]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > arrx[i] {
				arrx[i] = grid[i][j]
			}
			if grid[i][j] > arry[j] {
				arry[j] = grid[i][j]
			}
			if grid[i][j] > 0 {
				sum++
			}
		}
	}
	for i := 0; i < n; i++ {
		sum += arrx[i] + arry[i]
	}
	return sum
}