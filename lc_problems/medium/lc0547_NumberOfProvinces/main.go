func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	parents := make([]int, n)
	for i := 1; i < n; i++ {
		parents[i] = i
	}

	var findParent func(i int) int
	findParent = func(i int) int {
		if parents[i] == i {
			return i
		}
		return findParent(parents[i])
	}

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if isConnected[i][j] == 0 {
				continue
			}
			pi := findParent(i)
			pj := findParent(j)
			if pi != pj {
				parents[pi] = pj
			}
		}
	}

	set := make(map[int]struct{})
	for _, p := range parents {
		index := findParent(p)
		if _, ok := set[index]; ok {
			continue
		}
		set[index] = struct{}{}
	}

	return len(set)
}