func singleNonDuplicate(nums []int) int {
	n := len(nums)
	for i := 1; i < n; i += 2 {
		if nums[i-1] != nums[i] {
			return nums[i-1]
		}
	}
	return nums[n-1]
}