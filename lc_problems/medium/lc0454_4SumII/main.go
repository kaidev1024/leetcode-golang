func fourSumCount(A []int, B []int, C []int, D []int) int {
	n := len(A)
	if len(A) == 0 {
		return 0
	}

	m := make(map[int]int, n*n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			m[C[i]+D[j]]++
		}
	}

	cnt := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			cnt += m[-(A[i] + B[j])]
		}
	}
	return cnt
}