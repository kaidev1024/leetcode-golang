func threeConsecutiveOdds(arr []int) bool {
	count := 0
	for _, v := range arr {
		if v%2 == 1 {
			count++
			if count == 3 {
				return true
			}
		} else {
			count = 0
		}
	}
	return false
}