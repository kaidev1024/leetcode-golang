func combinationSum2(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	n := len(candidates)
	sort.Ints(candidates)
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
		for i := start; i <= end; {
			v := candidates[i]
			arr = append(arr, v)
			helper(i+1, end, target-v, arr)
			arr = arr[:len(arr)-1]
			i++
			for i <= end && candidates[i] == candidates[i-1] {
				i++
			}
		}
	}

	helper(0, n-1, target, []int{})
	return result
}