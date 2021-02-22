func helper(c byte) int {
	if c == 'A' {
		return 0
	}
	if c == 'C' {
		return 1
	}
	if c == 'G' {
		return 2
	}
	return 3
}

func findRepeatedDnaSequences(s string) []string {
	mask := 1<<20 - 1

	m := make(map[int]int)
	n := len(s)
	if n < 10 {
		return nil
	}

	result := make([]string, 0)
	v := 0
	for i := 0; i < 9; i++ {
		v <<= 2
		v += helper(s[i])
	}
	for i := 9; i < n; i++ {
		v <<= 2
		v += helper(s[i])
		v &= mask
		if count, ok := m[v]; ok {
			if count == 1 {
				result = append(result, s[i-9:i+1])
			}

		}
		m[v]++

	}
	return result
}