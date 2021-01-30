func toUpper(b byte) byte {
	if b >= 'a' && b <= 'z' {
		return b - 'a' + 'A'
	}
	return b
}

func licenseKeyFormatting(S string, K int) string {
	l := len(S)
	newL := l + l/K

	result := make([]byte, newL)
	pos := newL - 1

	j := 0

	for i := l - 1; i >= 0; i-- {
		if S[i] != '-' {
			result[pos] = toUpper(S[i])
			pos--
			j++
			if j == K {
				j = 0
				result[pos] = '-'
				pos--
			}
		}
	}
	pos++
	if pos == newL {
		return ""
	}
	if result[pos] == '-' {
		pos++
	}

	return string(result[pos:])
}