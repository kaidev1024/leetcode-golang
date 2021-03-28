func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
	m := make(map[int][]bool)
	for i, ns := 0, len(reservedSeats); i < ns; i++ {
		seat := reservedSeats[i]
		if seat[1] == 1 || seat[1] == 10 {
			continue
		}
		if _, ok := m[seat[0]]; !ok {
			m[seat[0]] = make([]bool, 4)
		}
		m[seat[0]][(seat[1]-2)>>1] = true
	}
	result := 0
	for _, v := range m {
		if !v[0] && !v[1] {
			result++
		}
		if !v[2] && !v[3] {
			result++
		}
		if v[0] && v[3] && !v[1] && !v[2] {
			result++
		}
	}

	return result + 2*(n-len(m))
}

