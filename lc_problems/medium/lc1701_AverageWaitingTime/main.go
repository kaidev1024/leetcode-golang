func averageWaitingTime(customers [][]int) float64 {
	var cur int
	var total int
	for _, v := range customers {
		if cur <= v[0] {
			cur = v[0] + v[1]
			total += v[1]
		} else {
			cur += v[1]
			total += cur - v[0]
		}
	}
	return float64(total) / float64(len(customers))

}