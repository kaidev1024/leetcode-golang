
func searchRange(nums []int, target int) []int {
	n := len(nums)
	if n == 0 {
		return []int{-1, -1}
	}
	l, r := 0, n-1

	findFirst := func(l, r int) int {
		for l < r {
			mid := (l + r) >> 1
			if nums[mid] < target {
				l = mid + 1
			} else {
				r = mid
			}
		}
		if nums[l] == target {
			return l
		}
		return -1
	}

	first := findFirst(l, r)
	if first == -1 {
		return []int{-1, -1}
	}

	findLast := func(l, r int) int {
		for l < r {
			mid := (l + r + 1) >> 1
			if nums[mid] > target {
				r = mid - 1
			} else {
				l = mid
			}
		}
		if nums[l] == target {
			return l
		}
		return -1
	}
	last := findLast(l, r)
	return []int{first, last}
}