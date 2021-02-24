func smallestRangeI(A []int, K int) int {
	n := len(A)
	if n == 1 {
		return 0
	}

	maxV := A[0] - K
	minV := A[0] + K

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 1; i < n; i++ {
		maxV = max(maxV, A[i]-K)
		minV = min(minV, A[i]+K)
	}

	if maxV < minV {
		return 0
	}
	return maxV - minV
}