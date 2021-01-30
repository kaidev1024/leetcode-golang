package leetcode

func islandPerimeter(grid [][]int) int {
	r := len(grid)
	c := len(grid[0])

	result := 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if grid[i][j] == 1 {
				result += 4
				if i-1 >= 0 && grid[i-1][j] == 1 {
					result--
				}
				if i+1 < r && grid[i+1][j] == 1 {
					result--
				}
				if j-1 >= 0 && grid[i][j-1] == 1 {
					result--
				}
				if j+1 < c && grid[i][j+1] == 1 {
					result--
				}
			}
		}
	}
	return result
}
