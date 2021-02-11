func calPoints(ops []string) int {
	n := len(ops)
	pos := 0
	scores := make([]int, n)

	for i := 0; i < n; i++ {
		switch ops[i] {
		case "+":
			scores[pos] = scores[pos-1] + scores[pos-2]
			pos++
		case "D":
			scores[pos] = scores[pos-1] << 1
			pos++
		case "C":
			pos--
		default:
			scores[pos], _ = strconv.Atoi(ops[i])
			pos++
		}
	}
	sum := 0
	for i := 0; i < pos; i++ {
		sum += scores[i]
	}
	return sum
}