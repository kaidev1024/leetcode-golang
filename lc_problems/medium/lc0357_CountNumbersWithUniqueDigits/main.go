func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	counts := make([]int, 11)
	counts[1] = 10
	for i := 2; i <= 10; i++ {
		result := 9
		cur := 9
		for j := 2; j <= i; j++ {
			result *= cur
			cur--
		}
		counts[i] = result
	}
	result := 0
	for i := 1; i <= n; i++ {
		result += counts[i]
	}
	return result
}