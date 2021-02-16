func convert(a, b byte) int {
	if a == '0' {
		return 0
	}
	if int(a-'0')*10+int(b-'0') > 26 {
		return 0
	}
	return 1
}
func numDecodings(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	count1 := 1
	count2 := 1
	for i := 1; i < n; i++ {
		sum := 0
		if s[i] != '0' {
			sum += count2
		}
		sum += count1 * convert(s[i-1], s[i])

		count1 = count2
		count2 = sum
	}
	return count2
}