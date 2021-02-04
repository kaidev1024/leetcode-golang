func restoreString(s string, indices []int) string {
	l := len(s)
	result := make([]byte, l)
	for i := 0; i < l; i++ {
		result[indices[i]] = s[i]
	}

	return string(result)
}