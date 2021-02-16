func subsets(nums []int) [][]int {
	n := len(nums)
	result := make([][]int, 0)

	var helper func(start int, arr []int)

	helper = func(start int, arr []int) {
		if start > n {
			return
		}

		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		result = append(result, arrCopy)

		for i := start; i < n; i++ {
			arr = append(arr, nums[i])
			helper(i+1, arr)
			arr = arr[:len(arr)-1]
		}
	}

	helper(0, []int{})

	return result
}