func maxScore(s string) int {
	l := len(s)
	zeros := make([]int, l)
	ones := make([]int, l)

	count := 0
	for i := 0; i < l; i++ {
		if s[i] == '0' {
			count++
		}
		zeros[i] = count
	}
	count = 0
	for i := l - 1; i >= 0; i-- {
		if s[i] == '1' {
			count++
		}
		ones[i] = count
	}

	maxVal := 0
	for i := 1; i < l; i++ {
		v := zeros[i-1] + ones[i]
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}