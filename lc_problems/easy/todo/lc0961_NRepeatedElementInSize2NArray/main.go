func repeatedNTimes(A []int) int {
	for i := 2; i < len(A); i++ {
		if A[i] == A[i-1] || A[i] == A[i-2] {
			return A[i]
		}
	}
	return A[0]
}