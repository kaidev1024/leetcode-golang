func makeGood(s string) string {
	l := len(s)
	result := make([]byte, l)

	pos := 0
	for i := 0; i < l; i++ {
		if pos == 0 {
			result[pos] = s[i]
			pos++
			continue
		}

		distance := int(s[i]) - int(result[pos-1])
		if distance == 32 || distance == -32 {
			pos--
			continue
		}
		result[pos] = s[i]
		pos++
	}
	return string(result[:pos])
}