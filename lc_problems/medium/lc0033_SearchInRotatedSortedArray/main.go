func helper(nums []int, l, r, target int) int {
	if r < l {
		return -1
	}
	for l < r {
		mid := (l + r) >> 1
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if nums[l] == target {
		return l
	}
	return -1
}

func search(nums []int, target int) int {
	n := len(nums)

	if nums[0] < nums[n-1] {
		return helper(nums, 0, n-1, target)
	}

	l, r := 0, n-1
	mid := 0
	for l < r {
		mid = (l + r) >> 1
		if nums[mid] > nums[mid+1] {
			l = mid
			break
		}
		if nums[mid] > nums[0] {
			l = mid + 1
		} else {
			r = mid
		}
	}

	if target >= nums[0] {
		return helper(nums, 0, l, target)
	}
	return helper(nums, l+1, n-1, target)
}