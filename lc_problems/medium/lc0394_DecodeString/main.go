func isAlpha(char byte) bool {
	return char <= 'z' && char >= 'a'
}

func isDigit(char byte) bool {
	return char <= '9' && char >= '0'
}

func decodeString(s string) string {
	pos := 0
	n := len(s)
	getInteger := func() int {
		result := 0
		for pos < n && isDigit(s[pos]) {
			result *= 10
			result += int(s[pos] - '0')
			pos++
		}
		return result
	}
	getString := func() string {
		var sb strings.Builder
		for pos < n && isAlpha(s[pos]) {
			sb.WriteByte(s[pos])
			pos++
		}
		return sb.String()
	}

	var helper func() string
	helper = func() string {
		var ret strings.Builder
		for pos < n {
			count := 1
			curStr := ""
			if isDigit(s[pos]) {
				count = getInteger()
			}
			if isAlpha(s[pos]) {
				curStr = getString()
				ret.WriteString(strings.Repeat(curStr, count))
				if pos == n {
					return ret.String()
				}
				if s[pos] == ']' {
					pos++
					return ret.String()
				}
				continue
			}
			pos++
			curStr = helper()
			ret.WriteString(strings.Repeat(curStr, count))
			if pos < n && s[pos] == ']' {
				pos++
				return ret.String()
			}

		}

		return ret.String()
	}

	return helper()
}
