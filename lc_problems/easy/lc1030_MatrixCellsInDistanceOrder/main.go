type coord struct {
	r int
	c int
}

var directions = []coord{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	visited := make(map[coord]struct{})
	res := make([][]int, 0, R*C)
	q := []coord{}

	q = append(q, coord{r0, c0})
	for len(q) != 0 {
		curr := q[0]
		q = q[1:]

		if _, ok := visited[curr]; ok {
			continue
		}

		visited[curr] = struct{}{}
		res = append(res, []int{curr.r, curr.c})

		for _, dir := range directions {
			newCoord := coord{curr.r + dir.r, curr.c + dir.c}
			if newCoord.r >= 0 && newCoord.r < R && newCoord.c >= 0 && newCoord.c < C {
				q = append(q, newCoord)
			}
		}
	}

	return res
}