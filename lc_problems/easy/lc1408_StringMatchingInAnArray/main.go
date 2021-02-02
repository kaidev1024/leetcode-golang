func stringMatching(words []string) []string {
	n := len(words)

	m := make(map[string]struct{})

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if strings.Contains(words[i], words[j]) {
				if _, ok := m[words[j]]; !ok {
					m[words[j]] = struct{}{}
				}
			}
			if strings.Contains(words[j], words[i]) {
				if _, ok := m[words[i]]; !ok {
					m[words[i]] = struct{}{}
				}
			}
		}
	}

	l := len(m)
	result := make([]string, 0, l)
	for key, _ := range m {
		result = append(result, key)
	}
	return result
}