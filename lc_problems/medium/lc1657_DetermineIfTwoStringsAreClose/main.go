func closeStrings(word1 string, word2 string) bool {
	n1 := len(word1)
	n2 := len(word2)
	if n1 != n2 {
		return false
	}

	counts1 := make([]int, 26)
	for _, c := range word1 {
		counts1[c-'a']++
	}
	counts2 := make([]int, 26)
	for _, c := range word2 {
		if counts1[c-'a'] == 0 {
			return false
		}
		counts2[c-'a']++
	}

	freqs := make(map[int]int)
	for _, v := range counts1 {
		freqs[v]++
	}
	for _, v := range counts2 {
		freqs[v]--
		if freqs[v] < 0 {
			return false
		}
	}

	return true
}
