func check(nums []int) bool {
	flag := false
	nums = append(nums, nums[0])
	for i, l := 1, len(nums); i < l; i++ {
		if nums[i-1] > nums[i] {
			if flag {
				return false
			}
			flag = true
		}
	}
	return true
}