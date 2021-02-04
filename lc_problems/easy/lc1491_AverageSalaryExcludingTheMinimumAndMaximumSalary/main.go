func isPathCrossing(path string) bool {
	arr := [2]int{0, 0}
	m := make(map[[2]int]struct{})
	m[[2]int{0, 0}] = struct{}{}
	for _, c := range path {
		if c == 'N' {
			arr[1]++
		} else if c == 'S' {
			arr[1]--
		} else if c == 'E' {
			arr[0]++
		} else {
			arr[0]--
		}
		if _, ok := m[arr]; ok {
			return true
		}
		m[arr] = struct{}{}
	}
	return false
}