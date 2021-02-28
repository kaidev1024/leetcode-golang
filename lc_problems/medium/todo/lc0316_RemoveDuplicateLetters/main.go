func removeDuplicateLetters(s string) string {
	lastIndex := make([]int, 26)
	seen := make([]bool, 26)
	n := len(s)
	for i := 0; i < n; i++ {
		lastIndex[s[i]-'a'] = i
	}

	result := make([]byte, 0, n)

	for i := 0; i < n; i++ {
		char := s[i]
		pos := char - 'a'
		if seen[pos] {
			continue
		}
		for len(result) > 0 {
			l := len(result)
			lastChar := result[l-1]
			if char < lastChar && lastIndex[lastChar-'a'] > i {
				result = result[:l-1]
				seen[lastChar-'a'] = false
				continue
			}
			break
		}
		result = append(result, char)
		seen[char-'a'] = true
	}
	return string(result)
}