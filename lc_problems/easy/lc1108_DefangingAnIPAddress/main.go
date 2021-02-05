func defangIPaddr(address string) string {
	l := len(address)
	result := make([]byte, l+6)

	for i, pos := 0, 0; i < l; i++ {
		if address[i] == '.' {
			result[pos] = '['
			result[pos+1] = '.'
			result[pos+2] = ']'
			pos += 3
		} else {
			result[pos] = address[i]
			pos++
		}
	}

	return string(result)
}