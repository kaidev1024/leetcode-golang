func helper(s string, start, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func validPalindrome(s string) bool {
	n := len(s)
	start := 0
	end := n - 1
	for start < end {
		if s[start] != s[end] {
			break
		}
		start++
		end--
	}
	if start >= end {
		return true
	}

	return helper(s, start+1, end) || helper(s, start, end-1)
}