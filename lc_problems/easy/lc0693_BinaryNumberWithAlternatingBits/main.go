func hasAlternatingBits(n int) bool {
	v := n ^ (n >> 1)
	return v&(v+1) == 0
}