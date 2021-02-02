func sortByBits(arr []int) []int {
	nbits := func(v int) int {
		n := 0
		for v > 0 {
			n++
			v &= v - 1
		}
		return n
	}

	sort.Slice(arr, func(i, j int) bool {
		ni := nbits(arr[i])
		nj := nbits(arr[j])
		if ni == nj {
			return arr[i] < arr[j]
		}
		return ni < nj
	})

	return arr
}