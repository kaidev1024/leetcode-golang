import "sort"

func avoidFlood(rains []int) []int {
	zeros := make([]int, 0)
	lakes := make(map[int]int)
	n := len(rains)
	result := make([]int, n)
	for i := 0; i < n; i++ {
		if rains[i] == 0 {
			zeros = append(zeros, i)
			continue
		}
		if ind, ok := lakes[rains[i]]; ok {
			ll := len(zeros)
			if ll == 0 {
				return nil
			}
			indToRemove := sort.SearchInts(zeros, ind)
			if indToRemove == len(zeros) {
				return nil
			}
			result[zeros[indToRemove]] = rains[i]
			zeros = append(zeros[:indToRemove], zeros[indToRemove+1:]...)
			result[i] = -1
			lakes[rains[i]] = i
		} else {
			lakes[rains[i]] = i
			result[i] = -1
		}
	}
	for _, v := range zeros {
		result[v] = 1
	}
	return result
}
