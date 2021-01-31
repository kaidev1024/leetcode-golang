func maxLengthBetweenEqualCharacters(s string) int {
	result := -1

	m := make(map[byte]int)

	for i, val := range []byte(s) {
		if ind, ok := m[val]; ok {
			distance := i - ind - 1
			if distance > result {
				result = distance
			}
		} else {
			m[val] = i
		}
	}
	return result
}