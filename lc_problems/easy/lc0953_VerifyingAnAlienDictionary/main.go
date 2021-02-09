func isAlienSorted(words []string, order string) bool {
	helper := func(word1, word2 string) bool {
		if word1 == word2 {
			return true
		}
		n1 := len(word1)
		n2 := len(word2)
		n := n1
		if n > n2 {
			n = n2
		}

		pos := 0
		for ; pos < n; pos++ {
			if word1[pos] != word2[pos] {
				break
			}
		}

		if pos == n {
			return n1 < n2
		}

		ind1 := strings.Index(order, string(word1[pos:pos+1]))
		ind2 := strings.Index(order, string(word2[pos:pos+1]))
		return ind1 < ind2
	}

	n := len(words)
	for i := 1; i < n; i++ {
		if !helper(words[i-1], words[i]) {
			return false
		}
	}
	return true
}