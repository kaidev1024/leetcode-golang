func numEquivDominoPairs(dominoes [][]int) int {
	m := make(map[string]int)

	l := len(dominoes)
	for i := 0; i < l; i++ {
		p := dominoes[i]
		if p[0] > p[1] {
			p[0], p[1] = p[1], p[0]
		}
		m[fmt.Sprintf("%v,%v", p[0], p[1])]++
	}

	result := 0
	for _, v := range m {
		result += (v - 1) * v >> 1
	}
	return result
}