func makeConnected(n int, connections [][]int) int {
	nc := len(connections)
	if nc < n-1 {
		return -1
	}

	parents := make([]int, n)
	for i := 0; i < n; i++ {
		parents[i] = i
	}

	var findParent func(i int) int
	findParent = func(i int) int {
		if parents[i] == i {
			return i
		}
		return findParent(parents[i])
	}

	union := func(i, j int) {
		parents[findParent(i)] = findParent(j)
	}

	for _, conn := range connections {
		union(conn[0], conn[1])
	}

	result := 0
	for i := 0; i < n; i++ {
		if parents[i] == i {
			result++
		}
	}
	return result - 1
}