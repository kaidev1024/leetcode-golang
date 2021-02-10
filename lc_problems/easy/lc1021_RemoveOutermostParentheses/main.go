func removeOuterParentheses(S string) string {
	n := len(S)
	if n == 0 {
		return ""
	}
	result := make([]byte, 0, n)

	count := 0
	start := 0
	end := 0

	for start < n && end < n {
		if S[end] == '(' {
			count++
		} else {
			count--
		}

		if count == 0 {
			fmt.Println(start, end)
			result = append(result, S[start+1:end]...)
			start = end + 1
			end = start
			continue
		}
		end++
	}

	return string(result)
}