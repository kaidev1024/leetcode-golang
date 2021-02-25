func longestWord(words []string) string {
	m := make(map[string]struct{})
	sort.Slice(words, func(i, j int) bool {
		li := len(words[i])
		lj := len(words[j])
		if li == lj {
			return words[i] < words[j]
		}
		return li < lj
	})

	r := 0
	rw := ""
	m[""] = struct{}{}
	for i, n := 0, len(words); i < n; i++ {
		word := words[i]
		lw := len(word)
		if _, ok := m[word[:lw-1]]; ok {
			m[word] = struct{}{}
			if lw > r {
				r = lw
				rw = word
			}
		}
	}
	return rw
}