func minOperations(n int) int {
	if n%2 == 1 {
		return (n - 1) * (n + 1) / 2 / 2
	}
	return n * ((n-2)/2 + 1) / 2
}