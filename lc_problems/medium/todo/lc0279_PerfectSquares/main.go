func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func numSquares(n int) int {
	num := make([]int, n+1)
	for i := 1; i*i <= n; i++ {
		num[i*i] = 1
	}

	for i := 1; i <= n; i++ {
		if num[i] == 0 {
			num[i] = i
		}

		for j := 1; j*j < i; j++ {
			num[i] = min(num[i], num[i-j*j]+num[j*j])
		}
	}
	return num[n]
}