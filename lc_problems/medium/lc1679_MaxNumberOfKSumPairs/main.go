func maxOperations(nums []int, k int) int {
	m := make(map[int]int)
	r := 0
	for _, num := range nums {
		target := k - num
		if v, ok := m[target]; ok && v > 0 {
			r++
			m[target]--
		} else {
			m[num]++
		}
	}
	return r
}