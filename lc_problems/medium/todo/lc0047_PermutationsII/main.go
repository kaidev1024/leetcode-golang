func permuteUnique(nums []int) [][]int {
	n := len(nums)
	if n == 0 {
		return nil
	}

	sort.Ints(nums)

	ret := make([][]int, 0)

	var helper func(nnums, arr []int)
	helper = func(nnums, arr []int) {
		if len(arr) == n {
			temp := make([]int, n)
			copy(temp, arr)
			ret = append(ret, temp)
			return
		}

		for i, l := 0, len(nnums); i < l; i++ {
			if i != 0 && nnums[i] == nnums[i-1] {
				continue
			}
			newNums := append(append([]int{}, nnums[:i]...), nnums[i+1:]...)
			helper(newNums, append(append([]int{}, arr...), nnums[i]))
		}
	}
	helper(nums, nil)
	return ret
}