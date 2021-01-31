func maxDepth(s string) int {
	result := 0
	count := 0
	for _, val := range s {
		if val == '(' {
			count++
			if count > result {
				result = count
			}
			continue
		}
		if val == ')' {
			count--
			continue
		}
	}
	return result
}