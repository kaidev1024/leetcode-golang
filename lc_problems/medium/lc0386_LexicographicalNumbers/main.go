func lexicalOrder(n int) []int {
	if n < 10 {
		result := make([]int, 0)
		for i := 1; i <= n; i++ {
			result = append(result, i)
		}
		return result
	}

	result := make([]int, 0, n)

	convert := func(s, i int) int {
		str := fmt.Sprintf("%v%v", s, i)
		num, _ := strconv.Atoi(str)
		return num
	}

	var helper func(s int)
	helper = func(s int) {
		for i := 0; i < 10; i++ {
			si := convert(s, i)
			if si > n {
				break
			}
			if si == n {
				result = append(result, si)
				break
			}
			result = append(result, si)
			helper(si)
		}
	}
	for i := 1; i <= 9; i++ {
		result = append(result, i)
		helper(i)
	}
	return result
}