
func findItinerary(tickets [][]string) []string {
	n := len(tickets)
	m := make(map[string][]string)
	visited := make(map[string][]bool)
	for _, ticket := range tickets {
		m[ticket[0]] = append(m[ticket[0]], ticket[1])
		visited[ticket[0]] = append(visited[ticket[0]], false)
	}
	for _, v := range m {
		sort.Strings(v)
	}

	result := make([]string, 0, n)
	result = append(result, "JFK")

	var dfs func(pos int) bool
	dfs = func(pos int) bool {
		if pos == n {
			return true
		}

		current := result[pos]
		for i, v := range visited[current] {
			if v {
				continue
			}
			visited[current][i] = true
			next := m[current][i]
			result = append(result, next)
			pos++
			if dfs(pos) {
				return true
			}
			visited[current][i] = false
			result = result[:pos]
			pos--

		}
		return false
	}
	dfs(0)
	return result
}