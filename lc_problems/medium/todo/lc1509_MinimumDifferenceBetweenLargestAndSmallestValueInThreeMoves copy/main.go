func minDifference(nums []int) int {
	s1, s2, s3, s4 := 1000000000, 1000000000, 1000000000, 1000000000
	l1, l2, l3, l4 := -1000000000, -1000000000, -1000000000, -1000000000
	n := len(nums)
	if n <= 4 {
		return 0
	}

	for i := 0; i < n; i++ {
		if nums[i] > l1 {
			l4 = l3
			l3 = l2
			l2 = l1
			l1 = nums[i]
		} else if nums[i] > l2 {
			l4 = l3
			l3 = l2
			l2 = nums[i]
		} else if nums[i] > l3 {
			l4 = l3
			l3 = nums[i]
		} else if nums[i] > l4 {
			l4 = nums[i]
		}
		if nums[i] < s1 {
			s4 = s3
			s3 = s2
			s2 = s1
			s1 = nums[i]
		} else if nums[i] < s2 {
			s4 = s3
			s3 = s2
			s2 = nums[i]
		} else if nums[i] < s3 {
			s4 = s3
			s3 = nums[i]
		} else if nums[i] < s4 {
			s4 = nums[i]
		}
	}

	d1 := l1 - s4
	d2 := l2 - s3
	d3 := l3 - s2
	d4 := l4 - s1

	return min(d1, min(d2, min(d3, d4)))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}