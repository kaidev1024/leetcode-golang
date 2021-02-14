func uniqueMorseRepresentations(words []string) int {
	morses := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	transformations := make(map[string]bool)
	for _, s := range words {
		transformation := ""
		runes := []rune(s)
		for _, r := range runes {
			index := r - 'a'
			transformation += morses[index]
		}
		transformations[transformation] = true
	}
	return len(transformations)
}