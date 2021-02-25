func findKthLargest(nums []int, k int) int {
	n := len(nums)
	return kSelect(nums, 0, n-1, k)
}

func kSelect(nums []int, left int, right int, k int) int {
	if left == right {
		return nums[left]
	}

	p := right
	l := left
	r := right - 1
	for l <= r {
		for l <= r && nums[l] >= nums[p] {
			l++
		}

		for l <= r && nums[r] < nums[p] {
			r--
		}

		if l <= r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}
	nums[l], nums[p] = nums[p], nums[l]
	p = l

	if p > k-1 {
		return kSelect(nums, left, p-1, k)
	} else if p < k-1 {
		return kSelect(nums, p+1, right, k)
	}

	return nums[p]
}