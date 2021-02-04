func countOdds(low int, high int) int {
	if low%2 == 1 {
		low--
	}
	if high%2 == 1 {
		high++
	}
	return (high - low) >> 1
}