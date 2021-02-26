func maximumBinaryString(binary string) string {

	n := len(binary)
	result := make([]byte, n)

	index0 := -1
	count0 := 0

	for i := 0; i < n; i++ {
		if binary[i] == '0' {
			count0++
			if index0 == -1 {
				index0 = i
			}
		}
	}

	if count0 < 2 {
		return binary
	}
	index0 += count0 - 1
	result[index0] = '0'
	for i := 0; i < n; i++ {
		if i != index0 {
			result[i] = '1'
		}
	}

	return string(result)

}