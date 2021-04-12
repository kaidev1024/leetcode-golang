func countBattleships(board [][]byte) int {
	nRow := len(board)
	nCol := len(board[0])

	result := 0
	for i := 0; i < nRow; i++ {
		for j := 0; j < nCol; j++ {
			if board[i][j] == '.' {
				continue
			}
			if i == nRow-1 && (j == nCol-1 || board[i][j+1] == '.') {
				result++
				continue
			}
			if j == nCol-1 && (i == nRow-1 || board[i+1][j] == '.') {
				result++
				continue
			}
			if i < nRow-1 && j < nCol-1 && board[i+1][j] == '.' && board[i][j+1] == '.' {
				result++
			}
		}
	}
	return result
}