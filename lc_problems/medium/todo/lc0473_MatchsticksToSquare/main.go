func makesquare(nums []int) bool {
	sum, maxN, n := 0, 0, len(nums)
	for _, num := range nums {
		sum += num
		maxN = max(maxN, num)
	}

	if sum%4 != 0 {
		return false
	}

	edgeLength := sum >> 2
	if maxN > edgeLength {
		return false
	}

	sort.Ints(nums)
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	var helper func(pos, target int) bool
	helper = func(pos, target int) bool {
		for i := pos; i < n; i++ {
			if nums[i] == 0 || target < nums[i] {
				continue
			}
			if target == nums[i] {
				nums[i] = 0
				return true
			}
			temp := nums[i]
			nums[i] = 0
			if helper(i+1, target-temp) {
				return true
			}
			nums[i] = temp
		}
		return false
	}

	for i := 0; i < 3; i++ {

		if !helper(0, edgeLength) {
			return false
		}
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}