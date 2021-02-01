func checkIfExist(arr []int) bool {
	m := make(map[int]struct{})

	for _, v := range arr {
		d := v << 1
		if _, ok := m[d]; ok {
			return true
		}
		if v%2 == 0 {
			half := v >> 1
			if _, ok := m[half]; ok {
				return true
			}
		}
		m[v] = struct{}{}
	}
	return false
}