func reformatNumber(number string) string {
	l := len(number)
	size := l/3 + l + 1
	byteArr := make([]byte, size)

	count := 0
	pos := 0

	for i := 0; i < l; i++ {
		if number[i] != ' ' && number[i] != '-' {
			byteArr[pos] = number[i]
			pos++
			count++
			if count == 3 {
				byteArr[pos] = '-'
				pos++
				count = 0
			}
		}
	}
	pos--
	if byteArr[pos] == '-' {
		pos--
	}
	if byteArr[pos-1] == '-' {
		byteArr[pos-2], byteArr[pos-1] = byteArr[pos-1], byteArr[pos-2]
	}
	return string(byteArr[:pos+1])
}