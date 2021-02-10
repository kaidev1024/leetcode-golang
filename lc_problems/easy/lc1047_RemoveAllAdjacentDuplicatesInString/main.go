func removeDuplicates(S string) string {
	n := len(S)
	arr := make([]byte, 0, n)
	pos := -1
	for i := 0; i < n; i++ {
		if pos == -1 || arr[pos] != S[i] {
			arr = append(arr, S[i])
			pos++
			continue
		}
		arr = arr[:pos]
		pos--
	}
	return string(arr[:pos+1])
}