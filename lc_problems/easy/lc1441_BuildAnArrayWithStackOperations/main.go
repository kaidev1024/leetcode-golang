func buildArray(target []int, n int) []string {
	l := len(target)
	cur := 1

	result := make([]string, 0, l*2)
	for i := 0; i < l; {
		if cur == target[i] {
			result = append(result, "Push")
			i++
		} else {
			result = append(result, "Push")
			result = append(result, "Pop")
		}
		cur++
	}

	return result
}