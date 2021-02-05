func findOcurrences(text string, first string, second string) []string {
	arr := strings.Fields(text)
	l := len(arr)

	result := make([]string, 0, l)
	for i := 0; i < l-2; i++ {
		if arr[i] != first {
			continue
		}
		if arr[i+1] == second {
			result = append(result, arr[i+2])
		}
	}
	return result
}