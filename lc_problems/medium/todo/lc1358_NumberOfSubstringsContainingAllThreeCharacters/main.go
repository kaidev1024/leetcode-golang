func numberOfSubstrings(s string) int {
	ca, cb, cc := 0, 0, 0
	i, j := 0, 0
	result := 0
	n := len(s)

	for j < n {
		if s[j] == 'a' {
			ca++
		} else if s[j] == 'b' {
			cb++
		} else {
			cc++
		}

		for ca > 0 && cb > 0 && cc > 0 {
			if s[i] == 'a' {
				ca--
			} else if s[i] == 'b' {
				cb--
			} else {
				cc--
			}
			i++
		}
		result += i
		j++
	}
	return result
}