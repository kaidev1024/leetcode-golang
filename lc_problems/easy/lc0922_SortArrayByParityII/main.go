func sortArrayByParityII(A []int) []int {
	indo := 1
	inde := 0
	n := len(A)
	for indo < n && inde < n {
		for ; indo < n && A[indo]%2 == 1; indo += 2 {
		}
		for ; inde < n && A[inde]%2 == 0; inde += 2 {
		}
		if indo >= n || inde >= n {
			break
		}
		A[indo], A[inde] = A[inde], A[indo]
		indo += 2
		inde += 2
	}
	return A
}