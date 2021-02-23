func bulbSwitch(n int) int {
	result := 0
	for i := 1; i*i <= n; i++ {
		result++
	}
	return result
}