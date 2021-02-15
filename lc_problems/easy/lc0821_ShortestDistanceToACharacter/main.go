func shortestToChar(s string, c byte) []int {
	n := len(s)
	result := make([]int, n)

	d := n
	for i := 0; i < n; i++ {
		if s[i] == c {
			d = 0
			result[i] = d
			continue
		}
		d++
		result[i] = d
	}

	d = n
	for i := n - 1; i >= 0; i-- {
		if s[i] == c {
			d = 0
			continue
		}
		d++
		if result[i] > d {
			result[i] = d
		}
	}

	return result
}