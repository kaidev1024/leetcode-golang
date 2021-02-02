func arrayRankTransform(arr []int) []int {
	arrCopy := make([]int, len(arr))
	copy(arrCopy, arr)

	sort.Ints(arrCopy)

	ranks := make(map[int]int)

	currentRank := 1
	for _, n := range arrCopy {
		if _, ok := ranks[n]; ok {
			continue
		}

		ranks[n] = currentRank
		currentRank++
	}

	res := make([]int, 0, len(arr))
	for _, n := range arr {
		res = append(res, ranks[n])
	}

	return res
}