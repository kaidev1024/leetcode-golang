func sortArrayByParity(A []int) []int {
	indo := 0
	inde := 0

	n := len(A)
	for indo < n && inde < n {
		for indo < n && A[indo]%2 == 0 {
			indo++
		}
		inde = indo
		for inde < n && A[inde]%2 == 1 {
			inde++
		}
		if inde >= n {
			break
		}
		A[inde], A[indo] = A[indo], A[inde]
		indo++
		inde++
	}
	return A
}