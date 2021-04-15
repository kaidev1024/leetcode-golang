func findMaxLength(nums []int) int {
	m := map[int]int{
		0: -1,
	}

	count := 0
	ret := 0

	for i, v := range nums {
		if v == 1 {
			count++
		} else {
			count--
		}
		if index, ok := m[count]; ok {
			ret = max(ret, i-index)
		} else {
			m[count] = i
		}
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}