func halvesAreAlike(s string) bool {
	vowels := "aeiouAEIOU"
	l := len(s)
	firstHalf := 0
	secondHalf := 0
	isVowel := func(c byte) bool {
		for i := 0; i < 10; i++ {
			if vowels[i] == c {
				return true
			}
		}
		return false
	}
	l2 := l / 2
	for i := 0; i < l2; i++ {
		if isVowel(s[i]) {
			firstHalf++
		}
	}
	for i := l2; i < l; i++ {
		if isVowel(s[i]) {
			secondHalf++
		}
	}
	return firstHalf == secondHalf
}