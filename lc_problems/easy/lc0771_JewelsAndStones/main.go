func numJewelsInStones(J string, S string) int {
	jewels := make(map[rune]bool)
	for _, j := range J {
		jewels[j] = true
	}
	count := 0
	for _, s := range S {
		if jewels[s] {
			count++
		}
	}
	return count
}