func generateTheString(n int) string {
	if n%2 == 0 {
		return "a" + strings.Repeat("b", n-1)
	}
	return strings.Repeat("a", n)
}