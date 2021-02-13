func isOneBitCharacter(bits []int) bool {
	n := len(bits)
	result := false

	for i := 0; i < n; {
		if bits[i] == 0 {
			result = true
			i++
		} else {
			result = false
			i += 2
		}
	}
	return result
}