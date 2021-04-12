func wiggleMaxLength(nums []int) int {
	up := true
	result := 2
	n := len(nums)
	if n < 2 {
		return n
	}
	start := 1
	for start < n {
		if nums[start] == nums[start-1] {
			start++
		} else {
			break
		}
	}

	if start == n {
		return 1
	}

	if nums[start] < nums[start-1] {
		up = false
	}
	for i := start; i < n; i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > nums[i-1] {
			if up {
				continue
			} else {
				up = true
				result++
			}
		} else {
			if up {
				up = false
				result++
			}
		}
	}
	return result
}