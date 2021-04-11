func findSubsequences(nums []int) [][]int {
	var ret [][]int
	n := len(nums)

	var helper func(pos int, cur []int)
	helper = func(pos int, cur []int) {
		lc := len(cur)
		if pos == n {
			if lc >= 2 {
				temp := make([]int, lc)
				copy(temp, cur)
				ret = append(ret, temp)
			}
			return
		}

		if lc == 0 || nums[pos] >= cur[lc-1] {
			helper(pos+1, append(cur, nums[pos]))
		}
		if lc == 0 || nums[pos] != cur[lc-1] {
			helper(pos+1, cur)
		}
	}
	helper(0, []int{})
	return ret
}