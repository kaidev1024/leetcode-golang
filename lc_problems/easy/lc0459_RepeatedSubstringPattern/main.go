func repeatedSubstringPattern(s string) bool {
	counts := make([]int, 26)
	for _, c := range s {
		counts[c-'a']++
	}
	n := len(s)
	if n == 1 {
		return false
	}
	minC := n
	for _, count := range counts {
		if count == 0 {
			continue
		}
		if minC > count {
			minC = count
		}
	}

	if minC == 1 {
		return false
	}
	if minC == n {
		return true
	}
	for i := 1; i < minC; i++ {
		if minC%i != 0 {
			continue
		}
		count := minC / i
		if n%count != 0 {
			continue
		}
		l := n / count
		pattern := s[:l]
		pos := l
		for pos < n {
			if s[pos:pos+l] != pattern {
				break
			}
			pos += l
		}
		if pos == n {
			return true
		}
	}
	return false
}