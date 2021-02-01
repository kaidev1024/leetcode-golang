func removePalindromeSub(s string) int {
	l := len(s)
	if l == 0 {
		return 0
	}

	left := 0
	right := l - 1
	for left < right {
		if s[left] == s[right] {
			left++
			right--
		} else {
			return 2
		}
	}

	return 1
}