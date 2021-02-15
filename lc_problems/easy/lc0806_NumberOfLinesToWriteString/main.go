func numberOfLines(widths []int, s string) []int {
	r := 1
	c := 0
	n := len(s)

	for i := 0; i < n; i++ {
		width := widths[s[i]-'a']
		c += width
		if c > 100 {
			r++
			c = width
		}
	}
	return []int{r, c}
}