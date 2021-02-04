func xorOperation(n int, start int) int {
	result := 0

	for i := 0; i < n; i++ {
		result ^= start
		start += 2
	}
	return result
}