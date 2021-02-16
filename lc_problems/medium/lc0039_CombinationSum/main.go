func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	n := len(candidates)

	var helper func(start, end, target int, arr []int)

	helper = func(start, end, target int, arr []int) {
		if target == 0 {
			arrcopy := make([]int, len(arr))
			copy(arrcopy, arr)
			result = append(result, arrcopy)
			return
		}
		if target < 0 {
			return
		}
		for i := start; i <= end; i++ {
			v := candidates[i]
			arr = append(arr, v)
			helper(i, end, target-v, arr)
			arr = arr[:len(arr)-1]
		}
	}

	helper(0, n-1, target, []int{})
	return result
}