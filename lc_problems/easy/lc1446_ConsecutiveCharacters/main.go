func maxPower(s string) int {
	l := len(s)

	var c byte = ' '
	count := 0
	result := -1

	for i := 0; i < l; i++ {
		if c != s[i] {
			if count > result {
				result = count

			}
			c = s[i]
			count = 1
			continue

		}
		if c == s[i] {
			count++
		}
	}
	if count > result {
		result = count
	}

	return result
}