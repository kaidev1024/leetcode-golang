func isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func reverseOnlyLetters(S string) string {
	arr := []byte(S)
	n := len(arr)
	l := 0
	r := n - 1
	for l < r {
		for l < r && !isAlpha(arr[l]) {
			l++
		}
		for l < r && !isAlpha(arr[r]) {
			r--
		}
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
	return string(arr)
}