func canArrange(arr []int, k int) bool {
	counts := make(map[int]int)
	for _, v := range arr {
		v %= k
		if v < 0 {
			v += k
		}
		vv := k - v
		if count, _ := counts[vv]; count == 0 {
			counts[v]++
		} else {
			counts[vv]--
		}
	}
	for key, v := range counts {
		if key%k != 0 && v != 0 {
			return false
		}
	}
	return true
}