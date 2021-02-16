func combine(n int, k int) [][]int {
	result := make([][]int, 0)

	var helper func(start, res int, arr []int)
	helper = func(start, res int, arr []int) {
		if res == 0 {
			arrCopy := make([]int, k)
			copy(arrCopy, arr)
			result = append(result, arrCopy)
			return
		}

		for i := start; i <= n-res+1; i++ {
			arr = append(arr, i)
			helper(i+1, res-1, arr)
			arr = arr[:len(arr)-1]
		}
	}

	helper(1, k, []int{})

	return result
}

