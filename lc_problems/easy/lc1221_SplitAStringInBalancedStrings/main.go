func balancedStringSplit(s string) int {
	c := 0
	result := 0
	for i := range s {
		if s[i] == 'L' {
			c++
		} else {
			c--
		}
		if c == 0 {
			result++
		}
	}
	return result
}