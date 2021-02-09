func isMonotonic(A []int) bool {
	n := len(A)
	if n <= 2 {
		return true
	}

	pos := 1
	for pos < n {
		if A[pos] == A[pos-1] {
			pos++
		} else {
			break
		}
	}

	if pos == n {
		return true
	}

	dif := A[pos] - A[pos-1]
	pos++

	if dif > 0 {
		for pos < n {
			if A[pos] < A[pos-1] {
				return false
			}
			pos++
		}
		return true
	}

	for pos < n {
		if A[pos] > A[pos-1] {
			return false
		}
		pos++
	}
	return true
}