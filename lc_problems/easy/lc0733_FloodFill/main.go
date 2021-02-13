func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	n := []int{0, -1}
	s := []int{0, 1}
	w := []int{-1, 0}
	e := []int{1, 0}
	dir := [][]int{n, s, w, e}

	oldColor := image[sr][sc]
	nr := len(image)
	nc := len(image[0])
	visited := make([][]bool, nr)
	for i := 0; i < nr; i++ {
		visited[i] = make([]bool, nc)
	}

	var helper func(sr, sc int)
	helper = func(sr, sc int) {

		if sr < 0 || sr >= nr || sc < 0 || sc >= nc {
			return
		}
		if visited[sr][sc] {
			return
		}
		if image[sr][sc] != oldColor {
			return
		}
		image[sr][sc] = newColor

		visited[sr][sc] = true
		for i := 0; i < 4; i++ {
			helper(sr+dir[i][0], sc+dir[i][1])
		}
	}
	helper(sr, sc)
	return image
}