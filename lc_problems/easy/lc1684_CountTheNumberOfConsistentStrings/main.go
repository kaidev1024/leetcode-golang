func countConsistentStrings(allowed string, words []string) int {
	counts := make([]bool, 128)
	for _, c := range allowed {
		counts[c] = true
	}

	lw := len(words)

	result := 0
	for i := 0; i < lw; i++ {

		ok := true
		for _, c := range words[i] {
			if !counts[c] {
				ok = false
				break
			}
		}
		if ok {
			result++
		}
	}
	return result
}