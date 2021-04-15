func updateMatrix(matrix [][]int) [][]int {
	offset := [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	m := len(matrix)
	n := len(matrix[0])

	res := make([][]int, m)
	var queue [][2]int
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, n)
		for j := 0; j < len(res[i]); j++ {
			if matrix[i][j] != 0 {
				res[i][j] = math.MaxInt64
			} else {
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	dis := 1
	for len(queue) != 0 {
		l := len(queue)
		for _, p := range queue {
			for _, dir := range offset {
				x := p[0] + dir[0]
				y := p[1] + dir[1]
				if x >= 0 && x < m && y >= 0 && y < n && res[x][y] > dis {
					res[x][y] = dis
					queue = append(queue, [2]int{x, y})
				}
			}
		}
		queue = queue[l:]
		dis++
	}
	return res
}