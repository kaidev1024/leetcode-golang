package leetcode

func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for _, v := range nums1 {
		m[v]++
	}

	result := make([]int, 0, len(nums2))
	for _, v := range nums2 {
		if count, ok := m[v]; ok && count > 0 {
			result = append(result, v)
			m[v]--
		}
	}
	return result
}
