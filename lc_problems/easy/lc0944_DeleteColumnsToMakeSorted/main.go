func minDeletionSize(A []string) int {
	var ret int
outer:
	for i := range A[0] {
		for j := 0; j < len(A)-1; j++ {
			if A[j][i] > A[j+1][i] {
				ret++
				continue outer
			}
		}
	}
	return ret
}