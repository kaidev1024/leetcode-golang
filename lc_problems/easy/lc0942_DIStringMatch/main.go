func diStringMatch(S string) []int {
	start := 0
	n := len(S)
	end := n
	result := make([]int, 0, end+1)
	for i := 0; i < n; i++ {
		if S[i] == 'I' {
			result = append(result, start)
			start++
		} else {
			result = append(result, end)
			end--
		}
	}
	result = append(result, start)
	return result
}