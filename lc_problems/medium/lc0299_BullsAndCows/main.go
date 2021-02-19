func getHint(secret string, guess string) string {
	digits := make([]int, 10)
	a := 0
	b := 0
	n := len(secret)
	for i := 0; i < n; i++ {
		if secret[i] == guess[i] {
			a++
			continue
		}

		spos := secret[i] - '0'
		if digits[spos] < 0 {
			b++
		}
		digits[spos]++

		gpos := guess[i] - '0'
		if digits[gpos] > 0 {
			b++
		}
		digits[gpos]--
	}

	return fmt.Sprintf("%vA%vB", a, b)
}