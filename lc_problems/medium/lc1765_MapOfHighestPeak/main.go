func highestPeak(isWater [][]int) [][]int {
	directions := make([][]int, 4)
	directions[0] = []int{0, 1}
	directions[1] = []int{0, -1}
	directions[2] = []int{1, 0}
	directions[3] = []int{-1, 0}

	nr := len(isWater)
	nc := len(isWater[0])
	fmt.Println(nr, nc)
	visited := make([][]bool, nr)
	for i := 0; i < nr; i++ {
		visited[i] = make([]bool, nc)
	}

	result := make([][]int, nr)
	for i := 0; i < nr; i++ {
		result[i] = make([]int, nc)
	}

	level := make([][]int, 0)
	for i := 0; i < nr; i++ {
		for j := 0; j < nc; j++ {
			if isWater[i][j] == 1 {
				level = append(level, []int{i, j})
				visited[i][j] = true
			}
		}
	}

	v := 0
	for len(level) > 0 {
		nextLevel := make([][]int, 0)
		for i, n := 0, len(level); i < n; i++ {
			cell := level[i]
			result[cell[0]][cell[1]] = v
			for j := 0; j < 4; j++ {
				r := cell[0] + directions[j][0]
				c := cell[1] + directions[j][1]
				if r < 0 || r >= nr || c < 0 || c >= nc {
					continue
				}
				if visited[r][c] {
					continue
				}
				visited[r][c] = true

				nextLevel = append(nextLevel, []int{r, c})
			}
		}
		v++
		level = nextLevel
	}
	return result
}