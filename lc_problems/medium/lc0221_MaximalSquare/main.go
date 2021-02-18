func maximalSquare(matrix [][]byte) int {

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	n := len(matrix)
	m := len(matrix[0])

	var result byte = '0'

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 || j == 0 {
				if matrix[i][j] > result {
					result = matrix[i][j]
				}
				continue
			}
			if matrix[i][j] == '1' {
				matrix[i][j] = min(matrix[i-1][j-1], matrix[i-1][j], matrix[i][j-1]) + 1
			}
			if matrix[i][j] > result {
				result = matrix[i][j]
			}
		}
	}
	return int(result-'0') * int(result-'0')
}

func min(a byte, b byte, c byte) byte {
	if b < a {
		a = b
	}
	if c < a {
		a = c
	}
	return a
}