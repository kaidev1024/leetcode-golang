func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) < 1 {
		return false
	}
	n := len(matrix[0])
	if n < 1 {
		return false
	}
	low, hi := 0, len(matrix)-1
	for low <= hi {
		m := (low + hi) / 2
		if matrix[m][0] <= target && target <= matrix[m][n-1] {
			colLow, colHi := 0, n-1
			for colLow <= colHi {
				colM := (colLow + colHi) / 2
				if target == matrix[m][colM] {
					return true
				}
				if target < matrix[m][colM] {
					colHi = colM - 1
				} else {
					colLow = colM + 1
				}

			}
		}
		if matrix[m][0] > target {
			hi = m - 1
		} else {
			low = m + 1
		}
	}

	return false
}