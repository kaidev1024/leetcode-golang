func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	nw1 := len(word1)
	nw2 := len(word2)

	pos1 := 0
	pos2 := 0

	w1 := 0
	w2 := 0

	lw1 := len(word1[0])
	lw2 := len(word2[0])

	for {
		if word1[w1][pos1] != word2[w2][pos2] {
			return false
		}
		pos1++
		if pos1 == lw1 {
			pos1 = 0
			w1++
			if w1 < nw1 {
				lw1 = len(word1[w1])
			}
		}
		pos2++
		if pos2 == lw2 {
			pos2 = 0
			w2++
			if w2 < nw2 {
				lw2 = len(word2[w2])
			}
		}
		if w1 == nw1 && w2 == nw2 {
			return true
		}
		if w1 == nw1 || w2 == nw2 {
			return false
		}
	}
	return true
}