func exist(board [][]byte, word string) bool {
	nr := len(board)
	nc := len(board[0])
	nw := len(word)

	var search func(r, c, w int) bool
	visited := make([][]bool, nr)
	for i := 0; i < nr; i++ {
		visited[i] = make([]bool, nc)
	}

	search = func(r, c, w int) bool {
		if visited[r][c] {
			return false
		}
		visited[r][c] = true
		if board[r][c] != word[w] {
			visited[r][c] = false
			return false
		}
		if w == nw-1 {
			return true
		}

		if r-1 >= 0 && search(r-1, c, w+1) {
			return true
		}
		if r+1 < nr && search(r+1, c, w+1) {
			return true
		}
		if c-1 >= 0 && search(r, c-1, w+1) {
			return true
		}
		if c+1 < nc && search(r, c+1, w+1) {
			return true
		}
		visited[r][c] = false
		return false
	}

	for i := 0; i < nr; i++ {
		for j := 0; j < nc; j++ {

			if search(i, j, 0) {
				return true
			}
		}
	}
	return false
}