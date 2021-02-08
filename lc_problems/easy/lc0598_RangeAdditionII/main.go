func maxCount(m int, n int, ops [][]int) int {
	minR := m
	minC := n
	l := len(ops)

	for i := 0; i < l; i++ {
		op := ops[i]
		if minR > op[0] {
			minR = op[0]
		}
		if minC > op[1] {
			minC = op[1]
		}
	}

	return minR * minC
}