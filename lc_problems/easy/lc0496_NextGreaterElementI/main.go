func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	n2 := len(nums2)
	stack := make([]int, n2)

	pos := -1
	for i := 0; i < n2; i++ {
		if pos == -1 {
			pos++
			stack[pos] = nums2[i]
		} else if nums2[i] < stack[pos] {
			pos++
			stack[pos] = nums2[i]
		} else {
			for pos >= 0 && nums2[i] > stack[pos] {
				m[stack[pos]] = nums2[i]
				pos--
			}
			pos++
			stack[pos] = nums2[i]
		}
	}

	for i, v := range nums1 {
		if val, ok := m[v]; ok {
			nums1[i] = val
		} else {
			nums1[i] = -1
		}
	}

	return nums1
}