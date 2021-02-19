

func nthUglyNumber(n int) int {
	if n <= 3 {
		return n
	}

	pos := []int{0, 0, 0}
	nums := []int{1}

	helper1 := func() int {
		i := nums[pos[0]]
		v := 2 * i
		i3 := nums[pos[1]]
		v3 := 3 * i3
		if v3 < v {
			v = v3
		}
		i5 := nums[pos[2]]
		v5 := 5 * i5
		if v5 < v {
			v = v5
		}
		return v
	}
	helper2 := func(v int) {
		vals := []int{2, 3, 5}
		for i := 0; i < 3; i++ {
			if vals[i]*(nums[pos[i]]) == v {
				pos[i]++
			}
		}
	}

	result := 1
	for i := 2; i <= n; i++ {
		result = helper1()
		helper2(result)
		nums = append(nums, result)
	}
	return result
}