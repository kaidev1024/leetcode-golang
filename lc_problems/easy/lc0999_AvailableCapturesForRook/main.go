func numRookCaptures(board [][]byte) int {
	count := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if string(board[i][j]) == "R" {
				return foundPawnFrom(i, j, board)
			}
		}
	}
	return count
}

func foundPawnFrom(i int, j int, board [][]byte) int {
	count := 0
	if i >= 0 && i <= len(board)-1 {
		tempI := i
		if i < len(board) {
			for tempI < len(board) {
				if string(board[tempI][j]) == "B" {
					break
				}
				if string(board[tempI][j]) == "p" {
					count++
					break
				}
				tempI++
			}
		}
		if i > 0 {
			tempI := i
			for tempI >= 0 {
				if string(board[tempI][j]) == "B" {
					break
				}
				if string(board[tempI][j]) == "p" {
					count++
					break
				}
				tempI--
			}
		}
	}

	if j >= 0 && j <= len(board[0]) {
		tempJ := j
		if j < len(board[0]) {
			for tempJ < len(board[0]) {
				if string(board[i][tempJ]) == "B" {
					break
				}
				if string(board[i][tempJ]) == "p" {
					count++
					break
				}
				tempJ++
			}
		}
		if j > 0 {
			tempJ := i
			for tempJ >= 0 {
				if string(board[i][tempJ]) == "B" {
					break
				}
				if string(board[i][tempJ]) == "p" {
					count++
					break
				}
				tempJ--
			}
		}
	}
	return count
}