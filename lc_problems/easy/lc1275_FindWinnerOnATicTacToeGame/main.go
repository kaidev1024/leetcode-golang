func tictactoe(moves [][]int) string {
	grid := make([][]byte, 3)
	for i := range grid {
		grid[i] = make([]byte, 3)
	}
	for i, v := range moves {
		if i%2 == 0 {
			grid[v[0]][v[1]] = 'X'
		} else {
			grid[v[0]][v[1]] = 'O'
		}
	}
	for i := 0; i < 3; i++ {
		if grid[i][0] == 'X' && grid[i][1] == 'X' && grid[i][2] == 'X' {
			return "A"
		}
		if grid[i][0] == 'O' && grid[i][1] == 'O' && grid[i][2] == 'O' {
			return "B"
		}
		if grid[0][i] == 'X' && grid[1][i] == 'X' && grid[2][i] == 'X' {
			return "A"
		}
		if grid[0][i] == 'O' && grid[1][i] == 'O' && grid[2][i] == 'O' {
			return "B"
		}
	}
	if grid[0][0] == 'X' && grid[1][1] == 'X' && grid[2][2] == 'X' {
		return "A"
	}
	if grid[0][0] == 'O' && grid[1][1] == 'O' && grid[2][2] == 'O' {
		return "B"
	}
	if grid[0][2] == 'X' && grid[1][1] == 'X' && grid[2][0] == 'X' {
		return "A"
	}
	if grid[0][2] == 'O' && grid[1][1] == 'O' && grid[2][0] == 'O' {
		return "B"
	}
	if len(moves) < 9 {
		return "Pending"
	}
	return "Draw"
}