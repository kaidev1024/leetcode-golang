func reduce(arr []int, n int) int {
	result := 0
	for i, l := 0, len(arr); i < l; i++ {
		result *= 10
		result += arr[i]
		result %= n
	}
	return result
}

func superPow(a int, b []int) int {
	a %= 1337
	if a == 0 || a == 1 {
		return a
	}

	nb := len(b)
	n := 1338

	if nb <= 4 {
		bint := 0
		for i := 0; i < nb; i++ {
			bint *= 10
			bint += b[i]
		}
		if bint < n {
			n = bint
		}
	}
	m := make(map[int]int)
	mr := make(map[int]int)
	m[0] = 1
	mr[1] = 0
	cur := a
	index := 1
	start := 0
	for index < n {
		m[cur] = index
		mr[index] = cur
		index++
		cur *= a
		cur %= 1337
		if v, ok := m[cur]; ok {
			start = v
			break
		}
	}
	if index == n {
		return cur
	}
	interval := index - start
	p := reduce(b, interval)
	return mr[p]
}