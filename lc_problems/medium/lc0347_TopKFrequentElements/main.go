func topKFrequent(nums []int, k int) []int {
	dict := make(map[int]int)
	for _, v := range nums {
		dict[v]++
	}

	var keys []int

	for k, _ := range dict {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i int, j int) bool {
		return dict[keys[i]] > dict[keys[j]]
	})

	return keys[:k]
}