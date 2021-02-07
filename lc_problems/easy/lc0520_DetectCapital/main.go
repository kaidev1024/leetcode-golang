
func detectCapitalUse(word string) bool {
	isCapital := func(c byte) bool {
		if c >= 'A' && c <= 'Z' {
			return true
		}
		return false
	}

	isAllCapital := func(arr string) bool {
		for _, v := range arr {
			if !isCapital(byte(v)) {
				return false
			}
		}
		return true
	}

	isAllSmall := func(arr string) bool {
		for _, v := range arr {
			if isCapital(byte(v)) {
				return false
			}
		}
		return true
	}

	if isCapital(word[0]) {
		return isAllCapital(word[1:]) || isAllSmall(word[1:])
	}
	return isAllSmall(word[1:])
}