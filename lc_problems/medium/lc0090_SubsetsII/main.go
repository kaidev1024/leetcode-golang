func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	result := make([][]int, 0)

	var helper func(start int, arr []int)

	helper = func(start int, arr []int) {
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		result = append(result, arrCopy)

		for i := start; i < n; i++ {
			if i != start && nums[i] == nums[i-1] {
				continue
			}
			arr = append(arr, nums[i])
			helper(i+1, arr)
			arr = arr[:len(arr)-1]
		}
	}
	helper(0, []int{})
	return result
}