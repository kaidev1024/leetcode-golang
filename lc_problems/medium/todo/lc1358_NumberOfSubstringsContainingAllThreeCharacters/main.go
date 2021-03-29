func numberOfSubstrings(s string) int {
	ca, cb, cc := 0, 0, 0
	result := 0
	n := len(s)

	for start, end := 0, 0; end < n; end++ {
		switch s[end] {
		case 'a':
			ca++
		case 'b':
			cb++
		case 'c':
			cc++
		}

		for ca > 0 && cb > 0 && cc > 0 {
			switch s[start] {
			case 'a':
				ca--
			case 'b':
				cb--
			case 'c':
				cc--
			}
			start++
		}

		result += start
	}
	return result
}