func combinationSum3(k int, n int) [][]int {
	result := make([][]int, 0)

	var helper func(start, k, n int, arr []int)
	helper = func(start, k, n int, arr []int) {
		if k == 0 {
			if n == 0 {
				arrCopy := make([]int, len(arr))
				copy(arrCopy, arr)
				result = append(result, arrCopy)
			}
			return
		}

		for i := start; i <= 9; i++ {
			arr = append(arr, i)
			helper(i+1, k-1, n-i, arr)
			arr = arr[:len(arr)-1]
		}
	}
	helper(1, k, n, []int{})
	return result
}