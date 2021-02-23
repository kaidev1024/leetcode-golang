func maxRepeating(sequence string, word string) int {
	ans, l1, l2 := 0, len(word), len(sequence)
	for i := 1; i*l1 <= l2; i++ {
		tmp := strings.Repeat(word, i)
		for j := 0; j+i*l1 <= l2; j++ {
			if sequence[j:j+i*l1] == tmp {
				ans = i
				break
			}
		}
	}
	return ans
}