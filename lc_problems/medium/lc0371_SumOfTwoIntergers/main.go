func getSum(a int, b int) int {
	if a == 0 || b == 0 {
		return a | b
	}
	return getSum(a^b, a&b<<1)
}