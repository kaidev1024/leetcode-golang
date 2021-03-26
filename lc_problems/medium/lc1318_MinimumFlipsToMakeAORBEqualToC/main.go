func minFlips(a int, b int, c int) int {
	count := 0
	for a|b != c {
		a0 := a & 1
		b0 := b & 1
		c0 := c & 1

		a >>= 1
		b >>= 1
		c >>= 1
		if a0|b0 == c0 {
			continue
		}
		if c0 == 1 {
			count++
			continue
		}
		if a0 == 1 {
			count++
		}
		if b0 == 1 {
			count++
		}
	}
	return count
}