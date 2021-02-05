func maxNumberOfBalloons(text string) int {
	balloon := make([]int, 128)
	balloon['b'] = 1
	balloon['a'] = 1
	balloon['l'] = 2
	balloon['o'] = 2
	balloon['n'] = 1

	counts := make([]int, 128)
	for _, v := range text {
		counts[v]++
	}

	b := "ablno"
	result := len(text)
	for i := 0; i < 5; i++ {
		c := counts[b[i]] / balloon[b[i]]
		if c < result {
			result = c
		}
	}
	return result
}