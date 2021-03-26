func minSteps(s string, t string) int {
	counts := make([]int, 26)
	for i, char := range s {
		counts[char-'a']++
		counts[t[i]-'a']--
	}

	n := 0
	for i := 0; i < 26; i++ {
		if counts[i] > 0 {
			n += counts[i]
		}
	}
	return n
}