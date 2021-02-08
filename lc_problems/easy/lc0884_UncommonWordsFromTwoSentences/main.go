func uncommonFromSentences(A string, B string) []string {
	wordsA := strings.Fields(A)
	wordsB := strings.Fields(B)

	counts := make(map[string]int)
	for i := 0; i < len(wordsA); i++ {
		counts[wordsA[i]]++
	}

	for i := 0; i < len(wordsB); i++ {
		counts[wordsB[i]]++
	}

	result := make([]string, 0, len(counts))
	for key, v := range counts {
		if v == 1 {
			result = append(result, key)
		}
	}
	return result
}