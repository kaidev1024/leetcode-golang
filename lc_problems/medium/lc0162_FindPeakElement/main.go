func findPeakElement(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}

	l, r := 0, n-1
	for l < r {
		mid := (l + r) / 2
		if mid == 0 {
			if nums[mid] > nums[mid+1] {
				return mid
			} else {
				l = mid + 1
			}
			continue
		}

		if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
			return mid
		}
		if nums[mid] > nums[mid-1] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}