type pair struct {
	key string
	val float64
}

func dfs(m map[string][]pair, start, end string, num float64, visited map[string]struct{}) float64 {
	if _, ok := visited[start]; ok {
		return -1.0
	}
	if start == end {
		return num
	}
	if start > end {
		return 1 / dfs(m, end, start, 1/num, visited)
	}
	visited[start] = struct{}{}
	if arr, ok := m[start]; ok {
		for _, v := range arr {
			temp := dfs(m, v.key, end, v.val*num, visited)
			if temp > 0.0 {
				return temp
			}
		}
	}
	delete(visited, start)
	return -1.0
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {

	m := make(map[string][]pair)
	set := make(map[string]struct{})
	for i, arr := range equations {
		if _, ok := set[arr[0]]; !ok {
			set[arr[0]] = struct{}{}
		}
		if _, ok := set[arr[1]]; !ok {
			set[arr[1]] = struct{}{}
		}
		m[arr[0]] = append(m[arr[0]], pair{arr[1], values[i]})
		m[arr[1]] = append(m[arr[1]], pair{arr[0], 1 / values[i]})
	}

	n := len(queries)
	result := make([]float64, n)
	for i, query := range queries {

		_, ok1 := set[query[0]]
		_, ok2 := set[query[1]]
		if !ok1 && !ok2 {
			result[i] = -1.0
			continue
		}
		result[i] = dfs(m, query[0], query[1], 1.0, make(map[string]struct{}))
	}
	return result
}