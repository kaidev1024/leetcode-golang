func sum(A []int) int {
	r := 0
	for _, v := range A {
		r += v
	}
	return r
}

func largestSumAfterKNegations(A []int, K int) int {
	sort.Ints(A)

	n := len(A)
	i := 0
	for ; K > 0 && i < n; i++ {
		if A[i] < 0 {
			A[i] = -A[i]
			K--
			continue
		}
		break
	}

	if K%2 == 0 {
		return sum(A)
	}

	if i == n {
		return sum(A) - 2*A[n-1]
	}

	if A[i] == 0 {
		return sum(A)
	}

	a := A[i]
	if i > 0 && a > A[i-1] {
		a = A[i-1]
	}
	return sum(A) - 2*a
}