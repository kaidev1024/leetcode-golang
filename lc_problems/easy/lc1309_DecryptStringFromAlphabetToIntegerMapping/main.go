func freqAlphabets(s string) string {
	l := len(s)

	result := make([]byte, l)
	pos := l - 1

	for i := l - 1; i >= 0; {
		if s[i] == '#' {
			num := (s[i-2]-'0')*10 + s[i-1] - '0'
			result[pos] = 'j' + num - 10
			pos--
			i -= 3
			continue
		}
		result[pos] = 'a' + (s[i] - '1')
		pos--
		i--
	}
	return string(result[pos+1:])
}