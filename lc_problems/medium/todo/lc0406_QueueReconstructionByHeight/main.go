func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		pi := people[i]
		pj := people[j]
		if pi[0] == pj[0] {
			return pi[1] > pj[1]
		}
		return pi[0] < pj[0]
	})

	n := len(people)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}

	result := make([][]int, n)
	for i := 0; i < n; i++ {
		p := people[i]
		index := p[1]
		result[arr[index]] = p
		arr = append(arr[:index], arr[index+1:]...)
	}
	return result
}