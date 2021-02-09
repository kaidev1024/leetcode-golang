func largestPerimeter(A []int) int {
	sort.Ints(A)
	n := len(A)

	for i := n - 1; i >= 2; i-- {
		for j := i - 1; j >= 1; j-- {
			for k := j - 1; k >= 0; k-- {
				if A[j]+A[k] > A[i] {
					return A[j] + A[k] + A[i]
				}
			}
		}
	}
	return 0
}