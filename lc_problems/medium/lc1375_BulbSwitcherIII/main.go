func numTimesAllBlue(light []int) int {
	dif := 0
	result := 0
	for i, v := range light {
		dif += v - i - 1
		if dif == 0 {
			result++
		}
	}
	return result
}