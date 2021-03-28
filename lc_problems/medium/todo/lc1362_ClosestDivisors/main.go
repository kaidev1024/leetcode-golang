func closestDivisors(num int) []int {
	r1 := calRes(num + 1)
	r2 := calRes(num + 2)
	if math.Abs(float64(r1[0]-r1[1])) > math.Abs(float64(r2[0]-r2[1])) {
		return r2
	}
	return r1
}

func calRes(i int) []int {
	sqrtOfI := int(math.Sqrt(float64(i)))
	for j := sqrtOfI; j > 0; j = j - 1 {
		if i%j == 0 {
			return []int{j, i / j}
		}
	}
	panic("no possibility")
}