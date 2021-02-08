func imageSmoother(M [][]int) [][]int {
	r := len(M)
	c := len(M[0])

	result := make([][]int, r)
	for i := 0; i < r; i++ {
		result[i] = make([]int, c)
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			sum := 0
			count := 0
			for m := -1; m <= 1; m++ {
				for n := -1; n <= 1; n++ {

					ii := i + m
					jj := j + n
					if ii < 0 || ii >= r || jj < 0 || jj >= c {
						continue
					}
					sum += M[ii][jj]
					count++
				}
			}
			result[i][j] = sum / count
		}
	}
	return result
}