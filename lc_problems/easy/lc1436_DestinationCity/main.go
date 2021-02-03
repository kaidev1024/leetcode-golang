func destCity(paths [][]string) string {
	m := make(map[string]struct{})

	for i, n := 0, len(paths); i < n; i++ {
		m[paths[i][0]] = struct{}{}
	}

	for i, n := 0, len(paths); i < n; i++ {
		if _, ok := m[paths[i][1]]; !ok {
			return paths[i][1]
		}
	}
	return ""
}