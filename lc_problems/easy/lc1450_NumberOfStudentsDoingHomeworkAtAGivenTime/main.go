func busyStudent(startTime []int, endTime []int, queryTime int) int {
	l := len(startTime)
	result := 0
	for i := 0; i < l; i++ {
		if queryTime >= startTime[i] && queryTime <= endTime[i] {
			result++
		}
	}
	return result
}