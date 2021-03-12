func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {
	n := len(source)
	parents := make([]int, n)
	for i := 0; i < n; i++ {
		parents[i] = i
	}

	find := func(p int) int {
		for {
			v := parents[p]
			if v == p {
				return p
			}
			p = v
		}
		return p
	}

	for _, pair := range allowedSwaps {
		p1 := find(pair[0])
		p2 := find(pair[1])
		parents[p1] = p2
	}

	groups := make(map[int][]int)

	for i := 0; i < n; i++ {
		groups[find(i)] = append(groups[find(i)], i)
	}

	ret := 0
	for _, arrInd := range groups {
		ni := len(arrInd)
		sourcei := make([]int, 0, ni)
		for _, ind := range arrInd {
			sourcei = append(sourcei, source[ind])
		}
		sort.Ints(sourcei)
		for _, ind := range arrInd {
			foundInd := sort.SearchInts(sourcei, target[ind])
			if foundInd >= len(sourcei) || sourcei[foundInd] != target[ind] {
				ret++
			} else {
				sourcei = append(sourcei[:foundInd], sourcei[foundInd+1:]...)
			}
		}
	}

	return ret
}
