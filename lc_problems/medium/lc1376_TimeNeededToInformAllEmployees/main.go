func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	m := make(map[int][]int)
	for i, v := range manager {
		if v == -1 {
			continue
		}
		m[v] = append(m[v], i)
	}

	ret := 0

	var dfs func(cur int, time int)
	dfs = func(cur int, time int) {
		if arr, ok := m[cur]; ok {
			for _, v := range arr {
				dfs(v, time+informTime[v])
			}
		} else {
			if time > ret {
				ret = time
			}
		}
	}

	dfs(headID, informTime[headID])
	return ret
}