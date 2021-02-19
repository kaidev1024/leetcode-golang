func hIndex(citations []int) int {
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})

	index := 0
	n := len(citations)
	for index < n {
		if citations[index] >= index+1 {
			index++
		} else {
			break
		}
	}
	return index
}