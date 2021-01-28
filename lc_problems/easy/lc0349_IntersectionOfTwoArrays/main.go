package leetcode

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]struct{})
	for _, v := range nums1 {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
		}
	}

	result := make([]int, 0, len(nums2))
	for _, v := range nums2 {
		if _, ok := m[v]; ok {
			result = append(result, v)
			delete(m, v)
		}
	}

	return result
}
