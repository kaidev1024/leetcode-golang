
func numMovesStones(a int, b int, c int) []int {
	if a > b {
		a, b = b, a
	}
	if b > c {
		b, c = c, b
		if a > b {
			a, b = b, a
		}
	}

	d1 := b - a
	d2 := c - b

	if d1 == 1 {
		if d2 == 1 {
			return []int{0, 0}
		}
		return []int{1, d2 - 1}
	}

	if d2 == 1 {
		if d1 == 1 {
			return []int{0, 0}
		}
		return []int{1, d1 - 1}
	}

	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	return []int{min(2, min(d1, d2)-1), d1 + d2 - 2}

}