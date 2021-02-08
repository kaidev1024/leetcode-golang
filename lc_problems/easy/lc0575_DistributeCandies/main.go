func distributeCandies(candyType []int) int {
	counts := make(map[int]struct{})
	l := len(candyType)
	for i := 0; i < l; i++ {
		counts[candyType[i]] = struct{}{}
	}

	n := len(counts)
	l >>= 1
	if n > l {
		return l
	}
	return n
}