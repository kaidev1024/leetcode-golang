func thousandSeparator(n int) string {
	s := strconv.Itoa(n)
	l := len(s)
	l2 := (l-1)/3 + l

	result := make([]byte, l2)

	count := 0
	j := l2 - 1

	for i := l - 1; i >= 0; i-- {
		result[j] = s[i]
		j--
		if j < 0 {
			break
		}
		count++
		if count == 3 {
			count = 0
			result[j] = '.'
			j--
		}
	}
	return string(result)
}