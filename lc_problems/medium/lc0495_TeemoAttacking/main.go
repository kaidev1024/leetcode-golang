func findPoisonedDuration(timeSeries []int, duration int) int {
	ret := 0
	endTime := 0
	for _, v := range timeSeries {
		if endTime <= v {
			ret += duration
		} else {
			ret += v + duration - endTime
		}
		endTime = v + duration
	}
	return ret
}