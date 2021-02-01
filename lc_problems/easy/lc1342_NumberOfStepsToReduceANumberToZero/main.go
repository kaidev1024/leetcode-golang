func numberOfSteps(num int) int {
	result := 0
	for num > 0 {
		if num%2 == 0 {
			num >>= 1
		} else {
			num--
		}
		result++
	}
	return result
}