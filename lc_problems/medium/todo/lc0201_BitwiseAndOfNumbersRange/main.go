func rangeBitwiseAnd(m int, n int) int {
	c := 0
	for {
		if m^n != 0 {
			m >>= 1
			n >>= 1
			c++
		} else {
			break
		}
	}
	return m << c
}