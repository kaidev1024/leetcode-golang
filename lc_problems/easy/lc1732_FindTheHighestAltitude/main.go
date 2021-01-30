func largestAltitude(gain []int) int {
	result := 0
	alt := 0
	for _, val := range gain {
		alt += val
		if result < alt {
			result = alt
		}
	}
	return result
}