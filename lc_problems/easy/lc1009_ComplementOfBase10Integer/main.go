func bitwiseComplement(N int) int {
	if N == 0 {
		return 1
	}
	result := 0
	pos := 0
	for N > 0 {
		cur := 1 - N&1
		fmt.Println(cur)
		result += cur << pos
		fmt.Println(result)
		pos++
		N >>= 1
	}

	return result
}