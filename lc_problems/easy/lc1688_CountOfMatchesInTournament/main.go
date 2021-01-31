func numberOfMatches(n int) int {
	result := 0
	for n > 1 {
		if n%2 == 0 {
			n /= 2
			result += n
		} else {
			n /= 2
			result += n
			n += 1
		}
	}
	return result
}