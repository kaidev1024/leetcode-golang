func decode(encoded []int, first int) []int {
	l := len(encoded)
	result := make([]int, l+1)
	result[0] = first
	for i := 0; i < l; i++ {
		result[i+1] = encoded[i] ^ result[i]
	}
	return result
}