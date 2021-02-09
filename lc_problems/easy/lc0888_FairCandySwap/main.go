func fairCandySwap(A []int, B []int) []int {
	ma := make(map[int]struct{})
	sumA := 0
	for _, v := range A {
		ma[v] = struct{}{}
		sumA += v
	}

	sumB := 0
	for _, v := range B {
		sumB += v
	}
	dif := (sumB - sumA) >> 1

	for _, v := range B {
		if _, ok := ma[v-dif]; ok {
			return []int{v - dif, v}
		}
	}
	return nil
}