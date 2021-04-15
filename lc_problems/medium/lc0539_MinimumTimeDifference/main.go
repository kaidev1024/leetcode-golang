func findMinDifference(timePoints []string) int {
	minutes := []int{}
	for _, timePoint := range timePoints {
		minutes = append(minutes, getMinutes(timePoint))
	}

	sort.Ints(minutes)
	min := math.MaxInt32
	for i := 1; i < len(minutes); i++ {
		if minutes[i]-minutes[i-1] < min {
			min = minutes[i] - minutes[i-1]
		}
	}

	first := minutes[0]
	last := minutes[len(minutes)-1]

	if 1440-last+first < min {
		return 1440 - last + first
	}

	return min
}

func getMinutes(timePoint string) int {
	t := strings.Split(timePoint, ":")
	tot := 0
	hours, _ := strconv.ParseInt(t[0], 10, 32)
	tot += int(hours) * 60
	minutes, _ := strconv.ParseInt(t[1], 10, 32)
	tot += int(minutes)

	return tot
}