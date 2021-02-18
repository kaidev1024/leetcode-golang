func findMin(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if nums[0] < nums[n-1] {
		return nums[0]
	}
	if nums[n-1] < nums[n-2] {
		return nums[n-1]
	}

	l, r := 0, n-1
	for l < r {
		mid := (l + r) >> 1
		if nums[mid] < nums[mid-1] && nums[mid] < nums[mid+1] {
			return nums[mid]
		}
		if nums[mid] > nums[n-1] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return nums[l]
}