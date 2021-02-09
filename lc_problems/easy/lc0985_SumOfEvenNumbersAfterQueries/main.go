func sumEvenAfterQueries(A []int, queries [][]int) []int {
	sum := 0
	n := len(A)
	for i := 0; i < n; i++ {
		if A[i]%2 == 0 {
			sum += A[i]
		}
	}

	nq := len(queries)
	result := make([]int, nq)

	for i := 0; i < nq; i++ {
		v := queries[i][0]
		ind := queries[i][1]
		if A[ind]%2 == 0 {
			sum -= A[ind]
		}
		A[ind] += v
		if A[ind]%2 == 0 {
			sum += A[ind]
		}
		result[i] = sum
	}
	return result
}