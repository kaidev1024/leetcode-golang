func relativeSortArray(arr1 []int, arr2 []int) []int {
	m := make(map[int]int)
	for i, v := range arr2 {
		m[v] = i
	}

	sort.Slice(arr1, func(i, j int) bool {
		v1, ok1 := m[arr1[i]]
		v2, ok2 := m[arr1[j]]

		if ok1 && ok2 {
			return v1 < v2
		}
		if !ok1 && !ok2 {
			return arr1[i] < arr1[j]
		}

		if !ok1 {
			return false
		}
		return true
	})
	return arr1
}