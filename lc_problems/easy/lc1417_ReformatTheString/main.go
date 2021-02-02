func reformat(s string) string {
	isDigit := func(v byte) bool {
		return v >= '0' && v <= '9'
	}
	isAlpha := func(v byte) bool {
		return v >= 'a' && v <= 'z'
	}

	nd := 0
	na := 0
	n := len(s)
	for i := 0; i < n; i++ {
		if isDigit(s[i]) {
			nd++
		} else {
			na++
		}
	}
	d := nd - na
	if d > 1 || d < -1 {
		return ""
	}

	result := make([]byte, n+1)

	id := 0
	ia := 0
	for i := 1; i < n && ia < n && id < n; i += 2 {
		for id < n {
			if !isDigit(s[id]) {
				id++
			} else {
				break
			}
		}
		for ia < n {
			if !isAlpha(s[ia]) {
				ia++
			} else {
				break
			}
		}
		result[i] = s[id]
		result[i+1] = s[ia]
		id++
		ia++
	}
	if nd == na {
		return string(result[1 : n+1])
	}

	if nd > na {
		for id < n {
			if !isDigit(s[id]) {
				id++
			} else {
				break
			}
		}
		result[n] = s[id]
		return string(result[1:])
	}
	for ia < n {
		if isDigit(s[ia]) {
			ia++
		} else {
			break
		}
	}
	result[0] = s[ia]
	return string(result[:n])
}