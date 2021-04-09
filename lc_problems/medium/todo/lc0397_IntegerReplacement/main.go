func integerReplacement(n int) int {
	ret := 0
	for n > 1 {
		if n%4 == 3 && n != 3 {
			n++
			ret++
			continue
		}
		if n%2 == 1 {
			n--
			ret++
			continue
		}
		n >>= 1
		ret++
	}
	return ret
}