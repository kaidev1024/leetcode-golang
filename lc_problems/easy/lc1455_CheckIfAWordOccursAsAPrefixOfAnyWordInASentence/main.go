func canBeEqual(target []int, arr []int) bool {
	counts := make([]int, 1001)
	for _, v := range target {
		counts[v]++
	}
	for _, v := range arr {
		counts[v]--
		if counts[v] < 0 {
			return false
		}
	}
	return true
}