func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})

	ret := 1
	cur := 0
	n := len(pairs)
	for i := 1; i < n; i++ {
		if pairs[cur][1] < pairs[i][0] {
			ret++
			cur = i
		}
	}
	return ret
}