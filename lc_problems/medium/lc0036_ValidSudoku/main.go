func isValidSudoku(board [][]byte) bool {
	validRow := func(r int) bool {
		filled := make([]bool, 9)
		for i := 0; i < 9; i++ {
			if board[r][i] == '.' {
				continue
			}
			pos := board[r][i] - '1'
			if filled[pos] {
				return false
			}
			filled[pos] = true
		}
		return true
	}
	validCol := func(c int) bool {
		filled := make([]bool, 9)
		for i := 0; i < 9; i++ {
			if board[i][c] == '.' {
				continue
			}
			pos := board[i][c] - '1'
			if filled[pos] {
				return false
			}
			filled[pos] = true
		}
		return true
	}
	validSquare := func(sr, er, sc, ec int) bool {
		filled := make([]bool, 9)
		for r := sr; r <= er; r++ {
			for c := sc; c <= ec; c++ {
				if board[r][c] == '.' {
					continue
				}
				pos := board[r][c] - '1'
				if filled[pos] {
					return false
				}
				filled[pos] = true
			}
		}
		return true
	}

	for i := 0; i < 9; i++ {
		if !validRow(i) {
			return false
		}
		if !validCol(i) {
			return false
		}
	}
	for r := 0; r < 9; r += 3 {
		for c := 0; c < 9; c += 3 {
			if !validSquare(r, r+2, c, c+2) {
				return false
			}
		}
	}
	return true
}